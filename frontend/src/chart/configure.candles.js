import { ExportCandleHistory } from "../../wailsjs/go/main/App";

import * as am5 from "@amcharts/amcharts5";
import * as am5xy from "@amcharts/amcharts5/xy";
import am5themes_Animated from "@amcharts/amcharts5/themes/Animated";
import { data } from './data/data.candles'

export async function loadCandleChart(def) {
  if (def) {
    return data
  } else {
    const candles = await ExportCandleHistory()
    candles.map((c) => {
      c.date = new Date(c.date).getTime();
      delete c.direction
    })
    return candles
  }
}

export function configureCandleChart(candleData, root, timeunit, timestep) {

    console.log('candles:', candleData)

    root.setThemes([
      am5themes_Animated.new(root)
    ]);

    var chart = root.container.children.push( 
      am5xy.XYChart.new(root, {
        panY: false,
        wheelY: "zoomX",
        layout: root.verticalLayout,
        maxtooltipDistance: 0
      }) 
    );
      // Create Y-axis

    var yRenderer = am5xy.AxisRendererY.new(root, {
      minGridDistance: 30,
      strokeOpacity: 0.17,
      stroke: am5.color(0xFFFFFF),
      strokeWidth: 1
    })

    yRenderer.labels.template.setAll({
      fill: am5.color(0xFFFFFF),
      fontSize: "0.71em",
      fillOpacity: 0.71
    })

    var yAxis = chart.yAxes.push(
      am5xy.ValueAxis.new(root, {
        renderer: yRenderer
      })
    );

    // Create X-Axis
    var xRenderer = am5xy.AxisRendererX.new(root, {
      minGridDistance: 50
    })

    xRenderer.labels.template.setAll({
      fill: am5.color(0xFFFFFF),
      fontSize: "0.71em",
      fillOpacity: 0.71
    })

    var xAxis = chart.xAxes.push(
      am5xy.DateAxis.new(root, {
        baseInterval: { timeUnit: `${timeunit}`, count: timestep },
        renderer: xRenderer
      })
    );

    // Create series
    var series = chart.series.push( 
      am5xy.CandlestickSeries.new(root, { 
        name: "Series",
        xAxis: xAxis, 
        yAxis: yAxis, 
        openValueYField: "open", 
        highValueYField: "high", 
        lowValueYField: "low", 
        valueYField: "close", 
        valueXField: "date",
        tooltip: am5.Tooltip.new(root, {})
      }) 
    );

    series.columns.template.states.create("riseFromOpen", {
      fill: am5.color(0x0B3600),
      stroke: am5.color(0x459630)
    });
    series.columns.template.states.create("dropFromOpen", {
      fill: am5.color(0x340006),
      stroke: am5.color(0x7B131F)
    });

    series.get("tooltip").label.set("text", "[bold]{valueX.formatDate()}[/]\nOpen: {openValueY}\nHigh: {highValueY}\nLow: {lowValueY}\nClose: {valueY}")
    series.data.setAll(candleData);


    // Add cursor
    chart.set("cursor", am5xy.XYCursor.new(root, {
      behavior: "zoomXY",
      xAxis: xAxis
    }));

    xAxis.set("tooltip", am5.Tooltip.new(root, {
      themeTags: ["axis"]
    }));

    yAxis.set("tooltip", am5.Tooltip.new(root, {
      themeTags: ["axis"]
    }));

}