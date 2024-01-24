//12306 book
function getRandomInterval(min, max) {
    // 生成指定范围内的随机整数
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function isOverReq(){ 
    let obj = document.querySelector("#query_ticket")
    if (obj && obj.className == "btn92s")
        return true
    return false
}
function clickSearch(){ 
    let obj = document.querySelector("#query_ticket")
    if (obj && obj.className == "btn92s")
        obj.click()
}
function clickBook(){ 
    let obj = document.querySelector("#ticket_5n0000K2300Q_03_14 > td.no-br > a")
    if (obj)
    {
    obj.click()
    console.log("booked")
    }  
}
function checkStateToBook(){
    let obj = document.querySelector("#YZ_5n0000K2300Q")
    if (obj && obj.innerText == "有")
        clickBook()
}
 
function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

isStop = false

async function task() {
    while (!isStop) {
        clickSearch();
       //  console.log("updateSearch")
        while (true) {
            await sleep(100);
            // 检查 isOverReq() 是否为 true
            if (isOverReq()) {
                break; // Exit the loop if condition is met
            }
        }
        console.log("bookCheck")
        checkStateToBook()
        // 等待100ms
        await sleep(getRandomInterval(3000, 5000));
    }
}
task()