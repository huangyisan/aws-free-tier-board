import dateformat from "dateformat";
export const chartType = "ColumnChart";

export const chartData = [

];

export const chartOptions = {
  chart: {
    title: "Trojan All Traffic",
    subtitle: ""
  },
  width: 800,
  height: 600
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