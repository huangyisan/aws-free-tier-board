import dateformat from "dateformat";
// import axios from "axios";

export const chartType = "ColumnChart";

export const chartData = [

];

export const chartOptions = {
  chart: {
    title: "Trojan All Traffic",
    subtitle: ""
  },
  // width: 1200,
  height: 400,
};


export var genAllTrafficData =  function(item, tmpdata) {
    let total = item.download + item.upload
    tmpdata.push([item.tag, item.download, item.upload, total ]);
}


export var genAllTrafficDataByGroup =  function(item, tmpdata) {
    let total = item.download + item.upload
    tmpdata.push([item.group, item.download, item.upload, total ]);
}

export var setDatesRange =  function(dates) {
    let today = new Date();
    today = dateformat(today, "yyyy-mm-dd");
    let tomorrow = new Date().setDate(new Date().getDate() + 1);
    tomorrow = dateformat(tomorrow, "yyyy-mm-dd");
    dates.push(today, tomorrow);
  }


export var genTrafficDataByTag = function(item, tmpdata) {
    let total = item.download + item.upload
    let theTime = item.time
    theTime = theTime.substr(0,16)
    tmpdata.push([theTime, item.download, item.upload, total ]);
}

export var genTrafficDataByGroup = function(item, tmpdata) {
  let total = item.download + item.upload
  let theTime = item.time
  theTime = theTime.substr(0,16)
  tmpdata.push([theTime, item.download, item.upload, total ]);
}