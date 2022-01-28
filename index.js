// 选择时间
function getTime() {
    var time = document.getElementById("timeSel");
    // var timecn = document.getElementsByClassName("timeCN");
    // var timeen = document.getElementById("timeEN");
    // var meal = document.getElementById("meal");
    return time;
}


// 选择饭店
function readJsonFile(file, callback) {
    var f = new XMLHttpRequest();
    f.overrideMimeType("application/json");
    f.open("GET", file, true);
    f.onreadystatechange = function() {
        if (f.readyState === 4 && f.status == "200") {
            callback(f.responseText);
        }
    }
    f.send(null)
}

function setRes(data) {
    console.log(data)
    // var select = document.getElementById(selectId);
    var resSel = document.getElementById("resSel")
    console.log(resSel)
    for (var i=0; i<data.length; i++) {
        var opt = document.createElement("option");
        opt.innerHTML = data[i].EN + " " + data[i].CN;
        resSel.appendChild(opt);
    }
}

readJsonFile("data.json", function(text){
    var data = JSON.parse(text);
    // console.log(data)
    // console.log(data[0].EN)
    setRes(data);
})

function getRes() {
    var resName = document.getElementById("resSel");
    return resName.value;
}

// 截单时间
function getEndTime() {
    var hour = document.getElementById("hour");
    var minute = document.getElementById("minute");
    var apm = document.getElementById("apm");
    return hour.value + ":" + minute.value + " " + apm.value;
}

function getSheetLink() {
    var sheetLink = document.getElementById("sheelLink");
    console.log(sheetLink)
    return sheetLink.value;
}

function getDrdLink() {
    var drdLink = document.getElementById("drdLink");
    return drdLink.value;
}

function generate() {
    var cn = document.getElementById("CN");
    var en = document.getElementById("EN");

    var time = getTime();
    var timecn, timeen, meal;
    if (time.value == "morning") {
        timecn = "中午";
        timeen = "Morning";
        meal = "Lunch";
    } else if (time.value == "evening") {
        timecn = "晚上";
        timeen = "Evening";
        meal = "Dinner";
    }
    var resName = getRes();
    var endtime = getEndTime();
    var sheetLink = getSheetLink();
    var drdLink = getDrdLink();

    cn.innerText = "哈喽大家" + timecn + "好 :blush: ，大家加班辛苦啦，我们" + timecn + "OT Meal目的地是：" + resName + "，需要上车点餐的老板请在下面的链接里sign up your full name在G列：（" + endtime + "截单）\n" + 
        sheetLink + "\n这是菜单链接: " + drdLink + "\nThanks and enjoy it！~:yum:";
    
    en.innerText = "Good " + timeen + " guys :blush: , we're going to order the OT meal for " + meal + " from: " + resName + ", please sign up your full name in column G below the OT Meal For " + meal + " if you need, here is the link for the name sheet:(" + 
        endtime + " end)\n" + sheetLink + "\nand menu:" + drdLink + "\nThanks and enjoy it! :yum:"
}