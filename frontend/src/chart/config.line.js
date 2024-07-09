import * as am5 from "@amcharts/amcharts5";
import * as am5xy from "@amcharts/amcharts5/xy";
import * as am5stock from "@amcharts/amcharts5/stock";
import am5themes_Animated from "@amcharts/amcharts5/themes/Animated";
import { data } from './data/data.line'

export function configureStockLine() {
    console.log("loaded'")
    let root = am5.Root.new("stockchart");

    root.setThemes([
      am5themes_Animated.new(root)
    ]);

   // Create a stock chart
    // https://www.amcharts.com/docs/v5/charts/stock-chart/#Instantiating_the_chart
    var stockChart = root.container.children.push(am5stock.StockChart.new(root, {
    }));

    /**
     * Main (value) panel
     */

    // Create a main stock panel (chart)
    // https://www.amcharts.com/docs/v5/charts/stock-chart/#Adding_panels
    var mainPanel = stockChart.panels.push(am5stock.StockPanel.new(root, {
    wheelY: "zoomX",
    panX: true,
    panY: true
    }));

    // Create axes
    // https://www.amcharts.com/docs/v5/charts/xy-chart/axes/
    var valueAxis = mainPanel.yAxes.push(am5xy.ValueAxis.new(root, {
    renderer: am5xy.AxisRendererY.new(root, {})
    }));

    var dateAxis = mainPanel.xAxes.push(am5xy.GaplessDateAxis.new(root, {
    baseInterval: {
        timeUnit: "day",
        count: 1
    },
    renderer: am5xy.AxisRendererX.new(root, {})
    }));

    // Add series
    // https://www.amcharts.com/docs/v5/charts/xy-chart/series/
    var valueSeries = mainPanel.series.push(am5xy.LineSeries.new(root, {
    name: "STCK",
    valueXField: "Date",
    valueYField: "Close",
    xAxis: dateAxis,
    yAxis: valueAxis,
    legendValueText: "{valueY}"
    }));

    valueSeries.data.setAll(data);

    // Set main value series
    // https://www.amcharts.com/docs/v5/charts/stock-chart/#Setting_main_series
    stockChart.set("stockSeries", valueSeries);

    // Add a stock legend
    // https://www.amcharts.com/docs/v5/charts/stock-chart/stock-legend/
    var valueLegend = mainPanel.plotContainer.children.push(am5stock.StockLegend.new(root, {
    stockChart: stockChart
    }));
    valueLegend.data.setAll([valueSeries]);

    /**
     * Secondary (volume) panel
     */

    // Create a main stock panel (chart)
    // https://www.amcharts.com/docs/v5/charts/stock-chart/#Adding_panels
    var volumePanel = stockChart.panels.push(am5stock.StockPanel.new(root, {
    wheelY: "zoomX",
    panX: true,
    panY: true,
    height: am5.percent(30)
    }));

    // Create axes
    // https://www.amcharts.com/docs/v5/charts/xy-chart/axes/
    var volumeValueAxis = volumePanel.yAxes.push(am5xy.ValueAxis.new(root, {
    numberFormat: "#.#a",
    renderer: am5xy.AxisRendererY.new(root, {})
    }));

    var volumeDateAxis = volumePanel.xAxes.push(am5xy.GaplessDateAxis.new(root, {
    baseInterval: {
        timeUnit: "day",
        count: 1
    },
    renderer: am5xy.AxisRendererX.new(root, {})
    }));

    // Add series
    // https://www.amcharts.com/docs/v5/charts/xy-chart/series/
    var volumeSeries = volumePanel.series.push(am5xy.ColumnSeries.new(root, {
    name: "STCK",
    valueXField: "Date",
    valueYField: "Volume",
    xAxis: volumeDateAxis,
    yAxis: volumeValueAxis,
    legendValueText: "{valueY}"
    }));

    volumeSeries.data.setAll(data);

    // Set main value series
    // https://www.amcharts.com/docs/v5/charts/stock-chart/#Setting_main_series
    stockChart.set("volumeSeries", volumeSeries);


    // Add a stock legend
    // https://www.amcharts.com/docs/v5/charts/stock-chart/stock-legend/
    var volumeLegend = volumePanel.plotContainer.children.push(am5stock.StockLegend.new(root, {
    stockChart: stockChart
    }));
    volumeLegend.data.setAll([volumeSeries]);


    // Add cursor(s)
    // https://www.amcharts.com/docs/v5/charts/xy-chart/cursor/
    mainPanel.set("cursor", am5xy.XYCursor.new(root, {
    yAxis: valueAxis,
    xAxis: dateAxis,
    snapToSeries: [valueSeries],
    snapToSeriesBy: "y!"
    }));

    volumePanel.set("cursor", am5xy.XYCursor.new(root, {
    yAxis: volumeValueAxis,
    xAxis: volumeDateAxis,
    snapToSeries: [volumeSeries],
    snapToSeriesBy: "y!"
    }));


    // Add scrollbar
    // https://www.amcharts.com/docs/v5/charts/xy-chart/scrollbars/
    var scrollbar = mainPanel.set("scrollbarX", am5xy.XYChartScrollbar.new(root, {
    orientation: "horizontal",
    height: 50
    }));
    stockChart.toolsContainer.children.push(scrollbar);

    var sbDateAxis = scrollbar.chart.xAxes.push(am5xy.GaplessDateAxis.new(root, {
    baseInterval: {
        timeUnit: "day",
        count: 1
    },
    renderer: am5xy.AxisRendererX.new(root, {})
    }));

    var sbValueAxis = scrollbar.chart.yAxes.push(am5xy.ValueAxis.new(root, {
    renderer: am5xy.AxisRendererY.new(root, {})
    }));

    var sbSeries = scrollbar.chart.series.push(am5xy.LineSeries.new(root, {
    valueYField: "Close",
    valueXField: "Date",
    xAxis: sbDateAxis,
    yAxis: sbValueAxis
    }));

    sbSeries.fills.template.setAll({
    visible: true,
    fillOpacity: 0.3
    });

    sbSeries.data.setAll(data);

    return root

}