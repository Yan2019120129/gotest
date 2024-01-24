
// 设置执行时间,开启定时任务
function executeAtSpecificTime(hours,min) {
    // 指定执行时间，这里假设是明天的 12:00:00
    const targetTime = new Date();
    targetTime.setHours(hours, min, 0, 0); // 设置时间为 12:00:00

    // 计算当前时间到执行时间的时间差
   // const timeDiff = targetTime.getTime() - new Date().getTime();
    Countdown(targetTime)
}

// 购票流程
function BuyTickets(){
    for (let i = 0; i < 100; i++) {
        clickBook()
        // 判断票数是否充足
        if(DetectionNum()) {
            // 票数充足 已经提交订单
           break
        }
        closeWin()
        console.log("重新请求订单")
    }
}

// 倒计时
function Countdown(targetTime){
    let intervalId=  setInterval(function () {
        var currentDate = new Date();
        var timeDifference = targetTime.getTime() - currentDate.getTime();

        // 计算倒计时的小时、分钟和秒
        const hours = Math.floor(timeDifference / (1000 * 60 * 60));
        const minutes = Math.floor((timeDifference % (1000 * 60 * 60)) / (1000 * 60));
        const seconds = Math.floor((timeDifference % (1000 * 60)) / 1000);
        console.log( "倒计时：" + hours + "小时 " + minutes + "分钟 " + seconds + "秒")
        // 如果目标日期已经过去，显示提示信息
        if (timeDifference === 0) {
            console.log("倒计时结束！")
            clearInterval(intervalId)
        }
    }, 1000);
}


// 执行发送请求
function clickBook(){
    let obj = document.querySelector("#toolbar_Div  div.btn.btn-hollow.btn-sm.w120.buy-ticket-button")
    // let obj = document.querySelector("#toolbar_Div > div.wrapper > div.ticket-result-box.ticket-fill-advance-result-box > div.ticket-result-bd > div > div.ticket-item-buy > div > div.ticket-item-buy-item > div.ticket-btn > a")
    console.log(obj)
    if (obj)  {
        console.log("开始购票")

        // 触发预购请求
        obj.click()

        // 提交订单
        // Confirm()
    }
}


// pop_170606405358314613
// 确认提交订单
function Confirm(){
    let obj = document.querySelector("body > div.dzp-confirm > div.modal > div.modal-ft > a.btn.btn-primary.ok")
    for (let i = 0; i <5; i++) {
        if (obj)  {
            obj.click()
            console.log("提交购票需求")
            break
        }
        obj = document.querySelector("body > div.dzp-confirm > div.modal > div.modal-ft > a.btn.btn-primary.ok")
    }
}



// #pop_17060665804324890 > div.modal > div.modal-bd > div > div.msg-con > h2
// 检测车票数量
// DetectionNum 检测车票数量
function DetectionNum(){
    let obj = document.querySelector("body > div.dzp-confirm > div.modal > div.modal-bd > div > div.msg-con > h2")
    if (obj){
       let message= obj.textContent.match(/\d+张/);
        // 输出匹配到的数字张
        if (message) {
            const numberString = message[0].match(/\d+/); // 再次使用正则表达式提取具体数字
            const number = parseInt(numberString, 10); // 将字符串转换成整数
            if (number !== 0){
                console.log("票数充足：" + number);
                return true;
            }
            console.log("票数不足：" + number);
            return false;
        } else {
            console.log("未找到匹配的数字。");
            return false;
        }
    }
    console.log("检测车票数量失败！！！");
    return false;
}

// #pop_17060665804324890 > div.modal > div.modal-ft > a
// 关闭窗口
function closeWin(){
    let obj = document.querySelector("body > div.dzp-confirm > div.modal > div.modal-ft > a")
    if (obj){
        obj.click()
        console.log("关闭票数显示窗口！！！")
    }
}

// 调用函数开始设置定时任务
executeAtSpecificTime(15,1);
