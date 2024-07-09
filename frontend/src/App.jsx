import { OpenStream, CloseStream } from "../wailsjs/go/main/App"

import { useEffect, useState, useRef } from 'react';
import './App.css';

import { configureCandleChart, loadCandleChart } from './chart/configure.candles';
import * as am5 from "@amcharts/amcharts5";

import Box from './box.jsx'

function App() {
    
    const chartDiv = useRef()
    const [streamCandles, setStreamCandles] = useState([])
    const [candleSource, setCandleSource] = useState(true)
    const [candleData, setCandleData] = useState([]) 

    const [autoGenRec, setAutoGenRec] = useState("single")
    const [autoGenActive, setAutoGenActive] = useState(false)
    const [intervalID, setIntervalID] = useState(null)

    const [autoGenMove, setAutoGenMove] = useState("up")
    const [slideValue, setSlideValue] = useState(50)
    
    const [selectedIndex, setSelectedIndex] = useState(null) 
    const [timeUnit, setTimeUnit] = useState("day")
    const [newCandleData, setNewCandleData] = useState({
        open: "",
        high: "",
        low: "",
        close: "",
    })

    const handleInput = (e, inputSource) => {
        console.log("HANDLE")
        const newValue = e.target.value
        newCandleData[inputSource] = newValue
        setNewCandleData({
            ...newCandleData 
        })
        console.log(newCandleData)
    }

    const [glyph, setGlyph] = useState([
        { level: 1000, state: -1, row: 1 },
        { level: 900, state: -1, row: 2  },
        { level: 800, state: -1, row: 3  },
        { level: 700, state: -1, row: 4  },
        { level: 600, state: -1, row: 5  },
        { level: 500, state: -1, row: 6  },
        { level: 400, state: -1, row: 7  },
        { level: 300, state: -1, row: 8  },
        { level: 200, state: -1, row: 9  },
        { level: 100, state: -1, row: 10  },
    ])

    function applyCSS(index) {
        if (glyph[index].state == 1) {
            return "up"
        } else {
            return "down"
        }
    }

    function updateState(index) {

        let old = glyph[index].state;
        let newval = old * -1
        
        glyph[index].state = newval
        setGlyph([...glyph])
    }

    function loopAutoGenerate(flag) {
        console.log("ACTIVE? ", autoGenActive)
        if (flag) {
            console.log('set interval creatined . . . ')
            setIntervalID(setInterval(autoGenerate, 1000))
        } else {
            console.log('clear the interval . . . .')
            clearInterval(intervalID)
        }
    }

    function autoGenerate() {
        console.log("autoGenerate fx called", autoGenMove)
        var newCandle = {}
        const lastCandle = candleData[candleData.length - 1]
        let delta 
        if(autoGenMove == "up") {
            delta = 0.009
            console.log('delta positive')
        } else {
            delta = -0.003
            console.log('delta negative')
        }
        for (var key in lastCandle) {
            if (key == "date") {
                const previousDate = new Date(lastCandle[key])
                previousDate.setTime(previousDate.getTime() + 1 * 86400000)
                newCandle[key] = previousDate.getTime()
            } else {
                
                let newValue = (lastCandle[key] + (lastCandle[key] * delta)).toFixed(2)
                
                if(autoGenMove == 'down' && key == 'close') {
                    newValue = (lastCandle['open'] - (lastCandle['open'] * -delta)).toFixed(2) - .02
                }

                newCandle[key] = parseFloat(newValue) 
            }
        }

        candleData.push(newCandle)
        setCandleData(candleData)
        
        chartDiv.current.dispose()
        chartDiv.current = am5.Root.new("candleschart"); 
        configureCandleChart(candleData, chartDiv.current, "day", 1)
    }

    function addNewCandle() {
        console.log("ADD NEW!")
        const length = candleData.length
        const lastCandle = candleData[length - 1]
        const newDate = new Date(lastCandle.date) 

        newDate.setTime(newDate.getTime() + 1 * 86400000)
        
        const newCandle = {
            date: newDate.getTime(),
            open: parseFloat(newCandleData.open),
            high: parseFloat(newCandleData.high),
            low: parseFloat(newCandleData.low),
            close: parseFloat(newCandleData.close)
        }

        candleData.push(newCandle)
        setCandleData(candleData)
        
        chartDiv.current.dispose()
        chartDiv.current = am5.Root.new("candleschart"); 
        configureCandleChart(candleData, chartDiv.current, "day", 1)
        
    }
    
    
    useEffect(() => {
        
        const candlesChart = am5.Root.new("candleschart"); 
        chartDiv.current = candlesChart

        console.log("HELLO!")
        const timeUnit = candleSource ? "day" : "minute"
        const timeStep = candleSource ? 1 : 5

        const fetchData = async () => {
            const candles = await loadCandleChart(candleSource)
            setCandleData(candles)
            configureCandleChart(candles, candlesChart, timeUnit, timeStep)
        }

        fetchData()       

        return () => {
            chartDiv.current.dispose()
        }

    }, [candleSource]);

    

    async function MockHistoryHandler() {
        const candleData = await OpenStream()
        const newestCandles = candleData

        console.log("wtfs",streamCandles)
        let copy = streamCandles

        newestCandles.forEach(newestCandle => {
            console.log("NewCandles", newestCandle.date)
    
            let newDate  = new Date(newestCandle.date)
            console.log('new date', newDate.getTime())
            const newCandle = {
                date: newDate.getTime(),
                open: parseFloat(newestCandle.open),
                high: parseFloat(newestCandle.high),
                low: parseFloat(newestCandle.low),
                close: parseFloat(newestCandle.close)
            }

            copy.push(newCandle)
        })
  
        setStreamCandles(copy)
        console.log("MY NEW HISTORY",copy)
        setCandleData(streamCandles)
        
        chartDiv.current.dispose()
        chartDiv.current = am5.Root.new("candleschart"); 
        configureCandleChart(streamCandles, chartDiv.current, "day", 1)
    }

    return (
        <div id="chartBody">
            <div className="panel">
                <div className="btnControls">
                    <div onClick={() => MockHistoryHandler()}>Real Candle Set</div>
                    {/* <div onClick={() => OpenStreamHandler()}>Open Stream</div>
                    <div onClick={() => CloseStream()}>Close Stream</div> */}
                </div>
            </div>
            <div className="display">
                <div id="candleschart" className="chartdiv"></div>
                {/* <div id="stockchart" className="chartdiv" style={{ width: "100%", height: "500px" }}></div> */}
                
                <div className="page-section">
                    <div className="controls">
                        <ul className="control-levels">
                            { 
                                glyph.map((el, i) => (
                                    <div className="control-row">
                                        <li key={`L-${el.level}`}>{el.level}</li>
                                        <li key={el.level} onClick={() => updateState(i)} className={applyCSS(i)}>{el.state}</li>
                                    </div>
                                )) 
                            }
                        </ul>
                        <div className="glyph">
                            { 
                                glyph.map((el) => (
                                    <Box up={el.state == 1 ? true : false} level={el.level} glyph={glyph} key={`B-${el.level}`}></Box>
                                )) 
                            }
                        </div> 
                    </div>
                </div>
            </div>
            <div className="panel">
                <div className="candle-info">
                    <ul className="candle-meta-view">
                        <li className="candle-meta header">
                            <span className="candle-date">TIMESTAMP</span>
                            <span className="candle-open">OPEN</span>
                            <span className="candle-high">HIGH</span>
                            <span className="candle-low">LOW</span>
                            <span className="candle-close">CLOSE</span>
                            <span className="candle-close">MOVED</span>
                        </li>
                        {
                            candleData.map((candle, index) => {
                                const date = new Date(candle.date)
                                const day = date.toLocaleDateString()
                                const time = date.toLocaleTimeString()
                                const timestamp = `${day} @ ${time}`
                                const direction = candle.open > candle.close ? "D" : "U"
                                const textColor = direction === "U" ? "#459630" : "#7B131F"
                                const opacity = selectedIndex === index ? 1 : 0.7
                                return (
                                <li className={`candle-meta ${selectedIndex === index ? 'selected' : ''}`} onClick={() => setSelectedIndex(index)}>
                                    <span className={`candle-date ${selectedIndex === index ? 'selected' : ''}`}>{timestamp}</span>
                                    <span className={`candle-open ${selectedIndex === index ? 'selected' : ''}`}>{candle.open}</span>
                                    <span className={`candle-high ${selectedIndex === index ? 'selected' : ''}`}>{candle.high}</span>
                                    <span className={`candle-low ${selectedIndex === index ? 'selected' : ''}`}>{candle.low}</span>
                                    <span className={`candle-close ${selectedIndex === index ? 'selected' : ''}`}>{candle.close}</span>
                                    <span className={`candle-direction ${selectedIndex === index ? 'selected' : ''}`} style={{color: textColor, opacity: opacity}}>{direction}</span>
                                </li>)})
                        }
                    </ul>
                    <div className="auto-candle-entry">
                        <div className="auto-candle-header">
                            <span>AutoGenerate:</span>
                            <div className="loop-control">
                                <span>Recurrence:</span>
                                <button className={`loop-ctrl-btn ${autoGenRec == "single" ? "selected" : ""}`} onClick={() => setAutoGenRec("single")}>Single</button>
                                <button className={`loop-ctrl-btn ${autoGenRec == "loop" ? "selected" : ""}`} onClick={() => setAutoGenRec("loop")}>Loop</button>
                            </div>
                        </div>
                        <div className="auto-interval">
                            <span>Interval:</span>
                            <div className="dropdown">
                                <select className="add-candle-timestamp">
                                    <option value="day">+ 1 day</option>
                                    <option value="M30">+ 30 min</option>
                                    <option value="S5">+ 5 sec</option>
                                </select>
                            </div>
                        </div>
                        <div className="auto-move-ctrl">
                            <span>Move:</span>
                            <div class="move-ctrl-btns">
                                <button className={`move-ctrl-btn-up ${autoGenMove == "up" ? "selected" : ""}`}
                                 onClick={() => {
                                    setAutoGenMove("up")
                                 }}>
                                    Up
                                </button>
                                <button className={`move-ctrl-btn-down ${autoGenMove == "down" ? "selected" : ""}`}
                                onClick={() => {
                                    setAutoGenMove("down")
                                }}>
                                    Down
                                </button>
                            </div>
                        </div>
                        <div class="slidecontainer">
                            <span>% Delta</span>
                            <div className="slidediv">
                                <input type="range" min="1" max="100" value={slideValue} class="slider" id="myRange" onChange={(e) => setSlideValue(e.target.value)} />
                                <span>{slideValue}</span>
                            </div>
                        </div>
                        <div className="init-autogen">
                        <button className={`loop-ctrl-btn ${autoGenActive ? "selected" : ""}`} onClick={() => {
                                const toggle = !autoGenActive
                                // console.log("SET ACTIVEATION FLAG TO:", toggle)
                                setAutoGenActive(toggle)
                                loopAutoGenerate(toggle)
                                // autoGenerate();
                            }}>Auto Generate</button>
                        </div>
                        
                    </div>
                    <div className="manual-controls">
                        <div className="manual-candle-entry">
                            <div className="manual-entry-header">
                                <span className="add-candle-title">Manual Entry:</span>
                                <div className="dropdown">
                                    <select className="add-candle-timestamp">
                                        <option value="day">+ 1 day</option>
                                        <option value="M30">+ 30 min</option>
                                        <option value="S5">+ 5 sec</option>
                                    </select>
                                </div>
                            </div>
                            <ul className="manual-entry-form">
                                <li className="meta-entry header">
                                </li>
                                <li className="meta-entry with-labels">
                                    <span className="candle-open">OPEN</span>
                                    <input placeholder="open" value={newCandleData.open} onChange={(e) => handleInput(e, "open")}></input>
                                    <span className="candle-high">HIGH</span>
                                    <input placeholder="high" value={newCandleData.high} onChange={(e) => handleInput(e, "high")}></input>
                                    <span className="candle-low">LOW</span>
                                    <input placeholder="low" value={newCandleData.low} onChange={(e) => handleInput(e, "low")}></input>
                                    <span className="candle-close">CLOSE</span>
                                    <input placeholder="close" value={newCandleData.close} onChange={(e) => handleInput(e, "close")}></input>
                                </li>
                            </ul>
                            <button className="add-candle" onClick={addNewCandle}>Add New Candle</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}
    
export default App
    