
// 设置执行时间,开启定时任务
function executeAtSpecificTime(hours,min) {
    // 指定执行时间，这里假设是明天的 12:00:00
    const targetTime = new Date();
    targetTime.setHours(hours, min, 0, 0); // 设置时间为 12:00:00

    // 计算当前时间到执行时间的时间差
   const timeDiff = targetTime.getTime() - new Date().getTime();
    console.log("在指定的时间执行的任务",timeDiff);

    setTimeout(function() {
        clickBook()
        // 这里可以调用你的任务函数或执行其他操作
    }, timeDiff);
}


// 执行发送请求
function clickBook(){
    let obj = document.querySelector("#toolbar_Div > div.wrapper > div.ticket-result-box.ticket-fill-advance-result-box > div.ticket-result-bd > div > div.ticket-item-buy > div > div.ticket-item-buy-item > div.ticket-btn > a")
    console.log("pre order:",obj)
    if (obj)  {
        console.log("开始购票")

        // 触发预购请求
        obj.click()

        // 提交订单
        Confirm()
    }
}


// pop_170606405358314613
// 确认提交订单
function Confirm(){
    let obj = document.querySelector("body > div.dzp-confirm > div.modal > div.modal-ft > a.btn.btn-primary.ok")
    while (true){
        console.log("Confirm:",obj)
        if (obj)  {
            obj.click()
            console.log("提交购票需求")
            break
        }
        obj = document.querySelector("body > div.dzp-confirm > div.modal > div.modal-ft > a.btn.btn-primary.ok")
    }
}

// 调用函数开始设置定时任务
executeAtSpecificTime(13,10);
