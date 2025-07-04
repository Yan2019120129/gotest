#!/bin/bash

# 硬盘健康与性能检测脚本
# 检测项目：SMART状态、坏块指标、读延迟、4K随机读IOPS
# 输出格式：每列一个硬盘，每行一个检测项目

# 检查必需工具
check_dependencies() {
    command -v smartctl &>/dev/null || { echo "请安装smartctl (smartmontools包)"; exit 1; }
    command -v ioping &>/dev/null || { echo "请安装ioping"; exit 1; }
    command -v fio &>/dev/null || { echo "请安装fio"; exit 1; }
    command -v bc &>/dev/null || { echo "请安装bc"; exit 1; }
}

# 获取硬盘列表
get_disks() {
    lsblk -d -n -o NAME,ROTA | awk '$2 == "1" {print "/dev/" $1}'
}

# 获取硬盘SN
get_disk_sn() {
    smartctl -i "$1" | awk -F': ' '/Serial Number:/ {print $2}' | tr -d ' '
}

# 检测SMART健康状态
check_smart_health() {
    smartctl -H "$1" | grep -q "PASSED" && echo "PASS" || echo "FAIL"
}

# 获取SMART属性值
get_smart_attr() {
    smartctl -A "$1" | awk -v id="$2" '$1 == id {print $10}'
}

# 测量读延迟
measure_latency() {
    local result
    result=$(ioping -c 10 -i 0 "$1" 2>/dev/null | awk '/avg/ {print $4}' | cut -d'/' -f2)
    [[ -n "$result" ]] && echo "$result" || echo "N/A"
}

# 测量4K随机读IOPS
measure_iops() {
    local output
    output=$(fio --name=test --filename="$1" --ioengine=libaio --direct=1 --bs=4k \
              --rw=randread --numjobs=1 --iodepth=1 --runtime=3 --time_based \
              --group_reporting --output-format=json 2>/dev/null)
    echo "$output" | grep -Eo '"iops": [0-9.]+' | awk '{print $2}' | head -1
}

# 格式化输出结果
format_result() {
    local value=$1 threshold=$2 type=$3
    case $type in
        "numeric")
            [[ -z "$value" || "$value" == "N/A" ]] && echo "N/A" && return
            if (( $(echo "$value < $threshold" | bc -l) )); then
                printf "%.2f (YES)" "$value"
            else
                printf "%.2f (NO)" "$value"
            fi
            ;;
        "count")
            [[ -z "$value" ]] && echo "N/A" && return
            if [[ "$value" -eq 0 ]]; then
                echo "0 (YES)"
            else
                echo "$value (NO)"
            fi
            ;;
        *)
            echo "$value"
            ;;
    esac
}

# 主函数
main() {
    check_dependencies
    
    declare -a disks
    mapfile -t disks < <(get_disks)
    
    [[ ${#disks[@]} -eq 0 ]] && { echo "未找到HDD硬盘"; exit 1; }
    
    # 收集数据
    declare -A sn smart_health reallocated pending uncorrectable latency iops
    
    for disk in "${disks[@]}"; do
        echo "[*] 正在检测 $disk ..."
        sn[$disk]=$(get_disk_sn "$disk")
        smart_health[$disk]=$(check_smart_health "$disk")
        reallocated[$disk]=$(get_smart_attr "$disk" "5")
        pending[$disk]=$(get_smart_attr "$disk" "197")
        uncorrectable[$disk]=$(get_smart_attr "$disk" "198")
        latency[$disk]=$(measure_latency "$disk")
        iops[$disk]=$(measure_iops "$disk")
    done
    
    # 打印结果表格
    echo -e "\n检测结果汇总:"
    printf "%15s" "硬盘"
    for disk in "${disks[@]}"; do
        printf "%20s" "$(basename "$disk")"
    done
    echo
    
    # SN行
    printf "%15s" "SN"
    for disk in "${disks[@]}"; do
        printf "%20s" "${sn[$disk]}"
    done
    echo
    
    # SMART健康状态
    printf "%15s" "SMART健康"
    for disk in "${disks[@]}"; do
        printf "%20s" "${smart_health[$disk]}"
    done
    echo
    
    # 05 重分配扇区
    printf "%15s" "05重分配扇区"
    for disk in "${disks[@]}"; do
        printf "%20s" "$(format_result "${reallocated[$disk]}" "1" "count")"
    done
    echo
    
    # C5 待映射扇区
    printf "%15s" "C5待映射扇区"
    for disk in "${disks[@]}"; do
        printf "%20s" "$(format_result "${pending[$disk]}" "1" "count")"
    done
    echo
    
    # C6 不可修复扇区
    printf "%15s" "C6不可修复扇区"
    for disk in "${disks[@]}"; do
        printf "%20s" "$(format_result "${uncorrectable[$disk]}" "1" "count")"
    done
    echo
    
    # 读延迟（ms）
    printf "%15s" "读延迟(ms)"
    for disk in "${disks[@]}"; do
        printf "%20s" "$(format_result "${latency[$disk]}" "50" "numeric")"
    done
    echo
    
    # 4K随机读IOPS
    printf "%15s" "4K随机读IOPS"
    for disk in "${disks[@]}"; do
        [[ "${iops[$disk]}" =~ ^[0-9.]+$ ]] && iops_val="${iops[$disk]}" || iops_val="N/A"
        if [[ "$iops_val" == "N/A" ]]; then
            printf "%20s" "N/A"
        else
            if (( $(echo "$iops_val > 100" | bc -l) )); then
                printf "%20s" "$(printf "%.0f (YES)" "$iops_val")"
            else
                printf "%20s" "$(printf "%.0f (NO)" "$iops_val")"
            fi
        fi
    done
    echo
}

main