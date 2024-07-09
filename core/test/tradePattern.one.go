package testData


// 1000U, 900D, 200U, 180D, 100U, 90D, trigger: 10U
//          90 ,  70,   88,   78,  87,

var trade_Pattern1_candles_forcedHigh = []RawCandle{
	{129.95, 130.00, 129.90, 129.95}, // high , i:0 //! HIGHEST
	{129.95, 130.00, 129.75, 129.80},
	{129.80, 130.00, 129.62, 129.70},
	{129.70, 130.00, 129.50, 129.52},
	{129.52, 130.00, 129.38, 129.41},
	{129.41, 130.00, 129.26, 129.28},
	{129.28, 130.00, 129.18, 129.21}, //TODO: removed conflicting .22 candle, reintroduce as test case
	{129.22, 130.00, 129.12, 129.16},
	{129.10, 130.00, 129.11, 129.11}, // 900 d:-1 i:8 //!
	{129.12, 130.00, 129.11, 129.18},
	{129.18, 130.00, 129.13, 129.15},
	{129.15, 130.00, 129.14, 129.20},
	{129.20, 130.00, 129.16, 129.24},
	{129.24, 130.00, 129.20, 129.28},
	{129.28, 130.00, 129.25, 129.30}, // 200 d:1 i:14
	{129.29, 130.00, 129.23, 129.24},
	{129.24, 130.00, 129.18, 129.19},
	{129.19, 130.00, 129.14, 129.15},
	{129.15, 130.00, 129.12, 129.12}, // 180 d:-1 i:18
	{129.12, 130.00, 129.12, 129.15},
	{129.15, 130.00, 129.10, 129.18}, //! .14 LOWEST
	{129.18, 130.00, 129.15, 129.22}, // 100 d:1 i:21
	{129.20, 130.00, 129.17, 129.16},
	{129.16, 130.00, 129.15, 129.14},
	{129.14, 130.00, 129.13, 129.13}, // 90 d:-1 i:24
}

var trade_Pattern1_candles_natural = []RawCandle{
	{129.95, 130.00, 129.90, 129.95}, // high , i:0
	{129.95, 130.85, 129.75, 129.80},
	{129.80, 130.76, 129.62, 129.70},
	{129.70, 130.72, 129.50, 129.52},
	{129.52, 130.56, 129.38, 129.41},
	{129.41, 130.44, 129.26, 129.28},
	{129.28, 130.32, 129.18, 129.21}, //TODO: removed conflicting .22 candle, reintroduce as test case
	{129.22, 130.24, 129.12, 129.16},
	{129.10, 130.18, 129.11, 129.11}, // 900 d:-1 i:8 //!
	{129.12, 130.20, 129.11, 129.18},
	{129.18, 130.22, 129.13, 129.15},
	{129.15, 130.22, 129.14, 129.20},
	{129.20, 130.28, 129.16, 129.24},
	{129.24, 130.29, 129.20, 129.28},
	{129.28, 130.31, 129.25, 129.30}, // 200 d:1 i:14
	{129.29, 130.28, 129.23, 129.24},
	{129.24, 130.22, 129.18, 129.19},
	{129.19, 130.18, 129.14, 129.15},
	{129.15, 130.14, 129.12, 129.12}, // 180 d:-1 i:18
	{129.12, 130.16, 129.12, 129.15},
	{129.15, 130.20, 129.10, 129.18}, //! .14
	{129.18, 130.22, 129.15, 129.22}, // 100 d:1 i:21
	{129.20, 130.22, 129.17, 129.16},
	{129.16, 130.17, 129.15, 129.14},
	{129.14, 130.14, 129.13, 129.13}, // 90 d:-1 i:24
}