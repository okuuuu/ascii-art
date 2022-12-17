package res

import "strings"

// Get color from its string determinant. Return white as default.

func GetColor(s string) int {
	colors := []struct {
		colorId int
		hex     string
		rgb     string
		hsl     string
		name    string
	}{
		{colorId: 0, hex: "#000000", rgb: "rgb(0,0,0)", hsl: "hsl(0%,0%,0%)", name: "black"},
		{colorId: 1, hex: "#800000", rgb: "rgb(128,0,0)", hsl: "hsl(0%,100%,25%)", name: "maroon"},
		{colorId: 2, hex: "#008000", rgb: "rgb(0,128,0)", hsl: "hsl(120%,100%,25%)", name: "green"},
		{colorId: 3, hex: "#808000", rgb: "rgb(128,128,0)", hsl: "hsl(60%,100%,25%)", name: "olive"},
		{colorId: 4, hex: "#000080", rgb: "rgb(0,0,128)", hsl: "hsl(240%,100%,25%)", name: "navy"},
		{colorId: 5, hex: "#800080", rgb: "rgb(128,0,128)", hsl: "hsl(300%,100%,25%)", name: "purple"},
		{colorId: 6, hex: "#008080", rgb: "rgb(0,128,128)", hsl: "hsl(180%,100%,25%)", name: "teal"},
		{colorId: 7, hex: "#c0c0c0", rgb: "rgb(192,192,192)", hsl: "hsl(0%,0%,75%)", name: "silver"},
		{colorId: 8, hex: "#808080", rgb: "rgb(128,128,128)", hsl: "hsl(0%,0%,50%)", name: "grey"},
		{colorId: 9, hex: "#ff0000", rgb: "rgb(255,0,0)", hsl: "hsl(0%,100%,50%)", name: "red"},
		{colorId: 10, hex: "#00ff00", rgb: "rgb(0,255,0)", hsl: "hsl(120%,100%,50%)", name: "lime"},
		{colorId: 11, hex: "#ffff00", rgb: "rgb(255,255,0)", hsl: "hsl(60%,100%,50%)", name: "yellow"},
		{colorId: 12, hex: "#0000ff", rgb: "rgb(0,0,255)", hsl: "hsl(240%,100%,50%)", name: "blue"},
		{colorId: 13, hex: "#ff00ff", rgb: "rgb(255,0,255)", hsl: "hsl(300%,100%,50%)", name: "fuchsia"},
		{colorId: 14, hex: "#00ffff", rgb: "rgb(0,255,255)", hsl: "hsl(180%,100%,50%)", name: "aqua"},
		{colorId: 15, hex: "#ffffff", rgb: "rgb(255,255,255)", hsl: "hsl(0%,0%,100%)", name: "white"},
		{colorId: 16, hex: "#000000", rgb: "rgb(0,0,0)", hsl: "hsl(0%,0%,0%)", name: "grey0"},
		{colorId: 17, hex: "#00005f", rgb: "rgb(0,0,95)", hsl: "hsl(240%,100%,18%)", name: "navyblue"},
		{colorId: 18, hex: "#000087", rgb: "rgb(0,0,135)", hsl: "hsl(240%,100%,26%)", name: "darkblue"},
		{colorId: 19, hex: "#0000af", rgb: "rgb(0,0,175)", hsl: "hsl(240%,100%,34%)", name: "blue3"},
		{colorId: 20, hex: "#0000d7", rgb: "rgb(0,0,215)", hsl: "hsl(240%,100%,42%)", name: "blue3"},
		{colorId: 21, hex: "#0000ff", rgb: "rgb(0,0,255)", hsl: "hsl(240%,100%,50%)", name: "blue1"},
		{colorId: 22, hex: "#005f00", rgb: "rgb(0,95,0)", hsl: "hsl(120%,100%,18%)", name: "darkgreen"},
		{colorId: 23, hex: "#005f5f", rgb: "rgb(0,95,95)", hsl: "hsl(180%,100%,18%)", name: "deepskyblue4"},
		{colorId: 24, hex: "#005f87", rgb: "rgb(0,95,135)", hsl: "hsl(197.777777777778%,100%,26%)", name: "deepskyblue4"},
		{colorId: 25, hex: "#005faf", rgb: "rgb(0,95,175)", hsl: "hsl(207.428571428571%,100%,34%)", name: "deepskyblue4"},
		{colorId: 26, hex: "#005fd7", rgb: "rgb(0,95,215)", hsl: "hsl(213.488372093023%,100%,42%)", name: "dodgerblue3"},
		{colorId: 27, hex: "#005fff", rgb: "rgb(0,95,255)", hsl: "hsl(217.647058823529%,100%,50%)", name: "dodgerblue2"},
		{colorId: 28, hex: "#008700", rgb: "rgb(0,135,0)", hsl: "hsl(120%,100%,26%)", name: "green4"},
		{colorId: 29, hex: "#00875f", rgb: "rgb(0,135,95)", hsl: "hsl(162.222222222222%,100%,26%)", name: "springgreen4"},
		{colorId: 30, hex: "#008787", rgb: "rgb(0,135,135)", hsl: "hsl(180%,100%,26%)", name: "turquoise4"},
		{colorId: 31, hex: "#0087af", rgb: "rgb(0,135,175)", hsl: "hsl(193.714285714286%,100%,34%)", name: "deepskyblue3"},
		{colorId: 32, hex: "#0087d7", rgb: "rgb(0,135,215)", hsl: "hsl(202.325581395349%,100%,42%)", name: "deepskyblue3"},
		{colorId: 33, hex: "#0087ff", rgb: "rgb(0,135,255)", hsl: "hsl(208.235294117647%,100%,50%)", name: "dodgerblue1"},
		{colorId: 34, hex: "#00af00", rgb: "rgb(0,175,0)", hsl: "hsl(120%,100%,34%)", name: "green3"},
		{colorId: 35, hex: "#00af5f", rgb: "rgb(0,175,95)", hsl: "hsl(152.571428571429%,100%,34%)", name: "springgreen3"},
		{colorId: 36, hex: "#00af87", rgb: "rgb(0,175,135)", hsl: "hsl(166.285714285714%,100%,34%)", name: "darkcyan"},
		{colorId: 37, hex: "#00afaf", rgb: "rgb(0,175,175)", hsl: "hsl(180%,100%,34%)", name: "lightseagreen"},
		{colorId: 38, hex: "#00afd7", rgb: "rgb(0,175,215)", hsl: "hsl(191.162790697674%,100%,42%)", name: "deepskyblue2"},
		{colorId: 39, hex: "#00afff", rgb: "rgb(0,175,255)", hsl: "hsl(198.823529411765%,100%,50%)", name: "deepskyblue1"},
		{colorId: 40, hex: "#00d700", rgb: "rgb(0,215,0)", hsl: "hsl(120%,100%,42%)", name: "green3"},
		{colorId: 41, hex: "#00d75f", rgb: "rgb(0,215,95)", hsl: "hsl(146.511627906977%,100%,42%)", name: "springgreen3"},
		{colorId: 42, hex: "#00d787", rgb: "rgb(0,215,135)", hsl: "hsl(157.674418604651%,100%,42%)", name: "springgreen2"},
		{colorId: 43, hex: "#00d7af", rgb: "rgb(0,215,175)", hsl: "hsl(168.837209302326%,100%,42%)", name: "cyan3"},
		{colorId: 44, hex: "#00d7d7", rgb: "rgb(0,215,215)", hsl: "hsl(180%,100%,42%)", name: "darkturquoise"},
		{colorId: 45, hex: "#00d7ff", rgb: "rgb(0,215,255)", hsl: "hsl(189.411764705882%,100%,50%)", name: "turquoise2"},
		{colorId: 46, hex: "#00ff00", rgb: "rgb(0,255,0)", hsl: "hsl(120%,100%,50%)", name: "green1"},
		{colorId: 47, hex: "#00ff5f", rgb: "rgb(0,255,95)", hsl: "hsl(142.352941176471%,100%,50%)", name: "springgreen2"},
		{colorId: 48, hex: "#00ff87", rgb: "rgb(0,255,135)", hsl: "hsl(151.764705882353%,100%,50%)", name: "springgreen1"},
		{colorId: 49, hex: "#00ffaf", rgb: "rgb(0,255,175)", hsl: "hsl(161.176470588235%,100%,50%)", name: "mediumspringgreen"},
		{colorId: 50, hex: "#00ffd7", rgb: "rgb(0,255,215)", hsl: "hsl(170.588235294118%,100%,50%)", name: "cyan2"},
		{colorId: 51, hex: "#00ffff", rgb: "rgb(0,255,255)", hsl: "hsl(180%,100%,50%)", name: "cyan1"},
		{colorId: 52, hex: "#5f0000", rgb: "rgb(95,0,0)", hsl: "hsl(0%,100%,18%)", name: "darkred"},
		{colorId: 53, hex: "#5f005f", rgb: "rgb(95,0,95)", hsl: "hsl(300%,100%,18%)", name: "deeppink4"},
		{colorId: 54, hex: "#5f0087", rgb: "rgb(95,0,135)", hsl: "hsl(282.222222222222%,100%,26%)", name: "purple4"},
		{colorId: 55, hex: "#5f00af", rgb: "rgb(95,0,175)", hsl: "hsl(272.571428571429%,100%,34%)", name: "purple4"},
		{colorId: 56, hex: "#5f00d7", rgb: "rgb(95,0,215)", hsl: "hsl(266.511627906977%,100%,42%)", name: "purple3"},
		{colorId: 57, hex: "#5f00ff", rgb: "rgb(95,0,255)", hsl: "hsl(262.352941176471%,100%,50%)", name: "blueviolet"},
		{colorId: 58, hex: "#5f5f00", rgb: "rgb(95,95,0)", hsl: "hsl(60%,100%,18%)", name: "orange4"},
		{colorId: 59, hex: "#5f5f5f", rgb: "rgb(95,95,95)", hsl: "hsl(0%,0%,37%)", name: "grey37"},
		{colorId: 60, hex: "#5f5f87", rgb: "rgb(95,95,135)", hsl: "hsl(240%,17%,45%)", name: "mediumpurple4"},
		{colorId: 61, hex: "#5f5faf", rgb: "rgb(95,95,175)", hsl: "hsl(240%,33%,52%)", name: "slateblue3"},
		{colorId: 62, hex: "#5f5fd7", rgb: "rgb(95,95,215)", hsl: "hsl(240%,60%,60%)", name: "slateblue3"},
		{colorId: 63, hex: "#5f5fff", rgb: "rgb(95,95,255)", hsl: "hsl(240%,100%,68%)", name: "royalblue1"},
		{colorId: 64, hex: "#5f8700", rgb: "rgb(95,135,0)", hsl: "hsl(77.7777777777778%,100%,26%)", name: "chartreuse4"},
		{colorId: 65, hex: "#5f875f", rgb: "rgb(95,135,95)", hsl: "hsl(120%,17%,45%)", name: "darkseagreen4"},
		{colorId: 66, hex: "#5f8787", rgb: "rgb(95,135,135)", hsl: "hsl(180%,17%,45%)", name: "paleturquoise4"},
		{colorId: 67, hex: "#5f87af", rgb: "rgb(95,135,175)", hsl: "hsl(210%,33%,52%)", name: "steelblue"},
		{colorId: 68, hex: "#5f87d7", rgb: "rgb(95,135,215)", hsl: "hsl(220%,60%,60%)", name: "steelblue3"},
		{colorId: 69, hex: "#5f87ff", rgb: "rgb(95,135,255)", hsl: "hsl(225%,100%,68%)", name: "cornflowerblue"},
		{colorId: 70, hex: "#5faf00", rgb: "rgb(95,175,0)", hsl: "hsl(87.4285714285714%,100%,34%)", name: "chartreuse3"},
		{colorId: 71, hex: "#5faf5f", rgb: "rgb(95,175,95)", hsl: "hsl(120%,33%,52%)", name: "darkseagreen4"},
		{colorId: 72, hex: "#5faf87", rgb: "rgb(95,175,135)", hsl: "hsl(150%,33%,52%)", name: "cadetblue"},
		{colorId: 73, hex: "#5fafaf", rgb: "rgb(95,175,175)", hsl: "hsl(180%,33%,52%)", name: "cadetblue"},
		{colorId: 74, hex: "#5fafd7", rgb: "rgb(95,175,215)", hsl: "hsl(200%,60%,60%)", name: "skyblue3"},
		{colorId: 75, hex: "#5fafff", rgb: "rgb(95,175,255)", hsl: "hsl(210%,100%,68%)", name: "steelblue1"},
		{colorId: 76, hex: "#5fd700", rgb: "rgb(95,215,0)", hsl: "hsl(93.4883720930233%,100%,42%)", name: "chartreuse3"},
		{colorId: 77, hex: "#5fd75f", rgb: "rgb(95,215,95)", hsl: "hsl(120%,60%,60%)", name: "palegreen3"},
		{colorId: 78, hex: "#5fd787", rgb: "rgb(95,215,135)", hsl: "hsl(140%,60%,60%)", name: "seagreen3"},
		{colorId: 79, hex: "#5fd7af", rgb: "rgb(95,215,175)", hsl: "hsl(160%,60%,60%)", name: "aquamarine3"},
		{colorId: 80, hex: "#5fd7d7", rgb: "rgb(95,215,215)", hsl: "hsl(180%,60%,60%)", name: "mediumturquoise"},
		{colorId: 81, hex: "#5fd7ff", rgb: "rgb(95,215,255)", hsl: "hsl(195%,100%,68%)", name: "steelblue1"},
		{colorId: 82, hex: "#5fff00", rgb: "rgb(95,255,0)", hsl: "hsl(97.6470588235294%,100%,50%)", name: "chartreuse2"},
		{colorId: 83, hex: "#5fff5f", rgb: "rgb(95,255,95)", hsl: "hsl(120%,100%,68%)", name: "seagreen2"},
		{colorId: 84, hex: "#5fff87", rgb: "rgb(95,255,135)", hsl: "hsl(135%,100%,68%)", name: "seagreen1"},
		{colorId: 85, hex: "#5fffaf", rgb: "rgb(95,255,175)", hsl: "hsl(150%,100%,68%)", name: "seagreen1"},
		{colorId: 86, hex: "#5fffd7", rgb: "rgb(95,255,215)", hsl: "hsl(165%,100%,68%)", name: "aquamarine1"},
		{colorId: 87, hex: "#5fffff", rgb: "rgb(95,255,255)", hsl: "hsl(180%,100%,68%)", name: "darkslategray2"},
		{colorId: 88, hex: "#870000", rgb: "rgb(135,0,0)", hsl: "hsl(0%,100%,26%)", name: "darkred"},
		{colorId: 89, hex: "#87005f", rgb: "rgb(135,0,95)", hsl: "hsl(317.777777777778%,100%,26%)", name: "deeppink4"},
		{colorId: 90, hex: "#870087", rgb: "rgb(135,0,135)", hsl: "hsl(300%,100%,26%)", name: "darkmagenta"},
		{colorId: 91, hex: "#8700af", rgb: "rgb(135,0,175)", hsl: "hsl(286.285714285714%,100%,34%)", name: "darkmagenta"},
		{colorId: 92, hex: "#8700d7", rgb: "rgb(135,0,215)", hsl: "hsl(277.674418604651%,100%,42%)", name: "darkviolet"},
		{colorId: 93, hex: "#8700ff", rgb: "rgb(135,0,255)", hsl: "hsl(271.764705882353%,100%,50%)", name: "purple"},
		{colorId: 94, hex: "#875f00", rgb: "rgb(135,95,0)", hsl: "hsl(42.2222222222222%,100%,26%)", name: "orange4"},
		{colorId: 95, hex: "#875f5f", rgb: "rgb(135,95,95)", hsl: "hsl(0%,17%,45%)", name: "lightpink4"},
		{colorId: 96, hex: "#875f87", rgb: "rgb(135,95,135)", hsl: "hsl(300%,17%,45%)", name: "plum4"},
		{colorId: 97, hex: "#875faf", rgb: "rgb(135,95,175)", hsl: "hsl(270%,33%,52%)", name: "mediumpurple3"},
		{colorId: 98, hex: "#875fd7", rgb: "rgb(135,95,215)", hsl: "hsl(260%,60%,60%)", name: "mediumpurple3"},
		{colorId: 99, hex: "#875fff", rgb: "rgb(135,95,255)", hsl: "hsl(255%,100%,68%)", name: "slateblue1"},
		{colorId: 100, hex: "#878700", rgb: "rgb(135,135,0)", hsl: "hsl(60%,100%,26%)", name: "yellow4"},
		{colorId: 101, hex: "#87875f", rgb: "rgb(135,135,95)", hsl: "hsl(60%,17%,45%)", name: "wheat4"},
		{colorId: 102, hex: "#878787", rgb: "rgb(135,135,135)", hsl: "hsl(0%,0%,52%)", name: "grey53"},
		{colorId: 103, hex: "#8787af", rgb: "rgb(135,135,175)", hsl: "hsl(240%,20%,60%)", name: "lightslategrey"},
		{colorId: 104, hex: "#8787d7", rgb: "rgb(135,135,215)", hsl: "hsl(240%,50%,68%)", name: "mediumpurple"},
		{colorId: 105, hex: "#8787ff", rgb: "rgb(135,135,255)", hsl: "hsl(240%,100%,76%)", name: "lightslateblue"},
		{colorId: 106, hex: "#87af00", rgb: "rgb(135,175,0)", hsl: "hsl(73.7142857142857%,100%,34%)", name: "yellow4"},
		{colorId: 107, hex: "#87af5f", rgb: "rgb(135,175,95)", hsl: "hsl(90%,33%,52%)", name: "darkolivegreen3"},
		{colorId: 108, hex: "#87af87", rgb: "rgb(135,175,135)", hsl: "hsl(120%,20%,60%)", name: "darkseagreen"},
		{colorId: 109, hex: "#87afaf", rgb: "rgb(135,175,175)", hsl: "hsl(180%,20%,60%)", name: "lightskyblue3"},
		{colorId: 110, hex: "#87afd7", rgb: "rgb(135,175,215)", hsl: "hsl(210%,50%,68%)", name: "lightskyblue3"},
		{colorId: 111, hex: "#87afff", rgb: "rgb(135,175,255)", hsl: "hsl(220%,100%,76%)", name: "skyblue2"},
		{colorId: 112, hex: "#87d700", rgb: "rgb(135,215,0)", hsl: "hsl(82.3255813953488%,100%,42%)", name: "chartreuse2"},
		{colorId: 113, hex: "#87d75f", rgb: "rgb(135,215,95)", hsl: "hsl(100%,60%,60%)", name: "darkolivegreen3"},
		{colorId: 114, hex: "#87d787", rgb: "rgb(135,215,135)", hsl: "hsl(120%,50%,68%)", name: "palegreen3"},
		{colorId: 115, hex: "#87d7af", rgb: "rgb(135,215,175)", hsl: "hsl(150%,50%,68%)", name: "darkseagreen3"},
		{colorId: 116, hex: "#87d7d7", rgb: "rgb(135,215,215)", hsl: "hsl(180%,50%,68%)", name: "darkslategray3"},
		{colorId: 117, hex: "#87d7ff", rgb: "rgb(135,215,255)", hsl: "hsl(200%,100%,76%)", name: "skyblue1"},
		{colorId: 118, hex: "#87ff00", rgb: "rgb(135,255,0)", hsl: "hsl(88.2352941176471%,100%,50%)", name: "chartreuse1"},
		{colorId: 119, hex: "#87ff5f", rgb: "rgb(135,255,95)", hsl: "hsl(105%,100%,68%)", name: "lightgreen"},
		{colorId: 120, hex: "#87ff87", rgb: "rgb(135,255,135)", hsl: "hsl(120%,100%,76%)", name: "lightgreen"},
		{colorId: 121, hex: "#87ffaf", rgb: "rgb(135,255,175)", hsl: "hsl(140%,100%,76%)", name: "palegreen1"},
		{colorId: 122, hex: "#87ffd7", rgb: "rgb(135,255,215)", hsl: "hsl(160%,100%,76%)", name: "aquamarine1"},
		{colorId: 123, hex: "#87ffff", rgb: "rgb(135,255,255)", hsl: "hsl(180%,100%,76%)", name: "darkslategray1"},
		{colorId: 124, hex: "#af0000", rgb: "rgb(175,0,0)", hsl: "hsl(0%,100%,34%)", name: "red3"},
		{colorId: 125, hex: "#af005f", rgb: "rgb(175,0,95)", hsl: "hsl(327.428571428571%,100%,34%)", name: "deeppink4"},
		{colorId: 126, hex: "#af0087", rgb: "rgb(175,0,135)", hsl: "hsl(313.714285714286%,100%,34%)", name: "mediumvioletred"},
		{colorId: 127, hex: "#af00af", rgb: "rgb(175,0,175)", hsl: "hsl(300%,100%,34%)", name: "magenta3"},
		{colorId: 128, hex: "#af00d7", rgb: "rgb(175,0,215)", hsl: "hsl(288.837209302326%,100%,42%)", name: "darkviolet"},
		{colorId: 129, hex: "#af00ff", rgb: "rgb(175,0,255)", hsl: "hsl(281.176470588235%,100%,50%)", name: "purple"},
		{colorId: 130, hex: "#af5f00", rgb: "rgb(175,95,0)", hsl: "hsl(32.5714285714286%,100%,34%)", name: "darkorange3"},
		{colorId: 131, hex: "#af5f5f", rgb: "rgb(175,95,95)", hsl: "hsl(0%,33%,52%)", name: "indianred"},
		{colorId: 132, hex: "#af5f87", rgb: "rgb(175,95,135)", hsl: "hsl(330%,33%,52%)", name: "hotpink3"},
		{colorId: 133, hex: "#af5faf", rgb: "rgb(175,95,175)", hsl: "hsl(300%,33%,52%)", name: "mediumorchid3"},
		{colorId: 134, hex: "#af5fd7", rgb: "rgb(175,95,215)", hsl: "hsl(280%,60%,60%)", name: "mediumorchid"},
		{colorId: 135, hex: "#af5fff", rgb: "rgb(175,95,255)", hsl: "hsl(270%,100%,68%)", name: "mediumpurple2"},
		{colorId: 136, hex: "#af8700", rgb: "rgb(175,135,0)", hsl: "hsl(46.2857142857143%,100%,34%)", name: "darkgoldenrod"},
		{colorId: 137, hex: "#af875f", rgb: "rgb(175,135,95)", hsl: "hsl(30%,33%,52%)", name: "lightsalmon3"},
		{colorId: 138, hex: "#af8787", rgb: "rgb(175,135,135)", hsl: "hsl(0%,20%,60%)", name: "rosybrown"},
		{colorId: 139, hex: "#af87af", rgb: "rgb(175,135,175)", hsl: "hsl(300%,20%,60%)", name: "grey63"},
		{colorId: 140, hex: "#af87d7", rgb: "rgb(175,135,215)", hsl: "hsl(270%,50%,68%)", name: "mediumpurple2"},
		{colorId: 141, hex: "#af87ff", rgb: "rgb(175,135,255)", hsl: "hsl(260%,100%,76%)", name: "mediumpurple1"},
		{colorId: 142, hex: "#afaf00", rgb: "rgb(175,175,0)", hsl: "hsl(60%,100%,34%)", name: "gold3"},
		{colorId: 143, hex: "#afaf5f", rgb: "rgb(175,175,95)", hsl: "hsl(60%,33%,52%)", name: "darkkhaki"},
		{colorId: 144, hex: "#afaf87", rgb: "rgb(175,175,135)", hsl: "hsl(60%,20%,60%)", name: "navajowhite3"},
		{colorId: 145, hex: "#afafaf", rgb: "rgb(175,175,175)", hsl: "hsl(0%,0%,68%)", name: "grey69"},
		{colorId: 146, hex: "#afafd7", rgb: "rgb(175,175,215)", hsl: "hsl(240%,33%,76%)", name: "lightsteelblue3"},
		{colorId: 147, hex: "#afafff", rgb: "rgb(175,175,255)", hsl: "hsl(240%,100%,84%)", name: "lightsteelblue"},
		{colorId: 148, hex: "#afd700", rgb: "rgb(175,215,0)", hsl: "hsl(71.1627906976744%,100%,42%)", name: "yellow3"},
		{colorId: 149, hex: "#afd75f", rgb: "rgb(175,215,95)", hsl: "hsl(80%,60%,60%)", name: "darkolivegreen3"},
		{colorId: 150, hex: "#afd787", rgb: "rgb(175,215,135)", hsl: "hsl(90%,50%,68%)", name: "darkseagreen3"},
		{colorId: 151, hex: "#afd7af", rgb: "rgb(175,215,175)", hsl: "hsl(120%,33%,76%)", name: "darkseagreen2"},
		{colorId: 152, hex: "#afd7d7", rgb: "rgb(175,215,215)", hsl: "hsl(180%,33%,76%)", name: "lightcyan3"},
		{colorId: 153, hex: "#afd7ff", rgb: "rgb(175,215,255)", hsl: "hsl(210%,100%,84%)", name: "lightskyblue1"},
		{colorId: 154, hex: "#afff00", rgb: "rgb(175,255,0)", hsl: "hsl(78.8235294117647%,100%,50%)", name: "greenyellow"},
		{colorId: 155, hex: "#afff5f", rgb: "rgb(175,255,95)", hsl: "hsl(90%,100%,68%)", name: "darkolivegreen2"},
		{colorId: 156, hex: "#afff87", rgb: "rgb(175,255,135)", hsl: "hsl(100%,100%,76%)", name: "palegreen1"},
		{colorId: 157, hex: "#afffaf", rgb: "rgb(175,255,175)", hsl: "hsl(120%,100%,84%)", name: "darkseagreen2"},
		{colorId: 158, hex: "#afffd7", rgb: "rgb(175,255,215)", hsl: "hsl(150%,100%,84%)", name: "darkseagreen1"},
		{colorId: 159, hex: "#afffff", rgb: "rgb(175,255,255)", hsl: "hsl(180%,100%,84%)", name: "paleturquoise1"},
		{colorId: 160, hex: "#d70000", rgb: "rgb(215,0,0)", hsl: "hsl(0%,100%,42%)", name: "red3"},
		{colorId: 161, hex: "#d7005f", rgb: "rgb(215,0,95)", hsl: "hsl(333.488372093023%,100%,42%)", name: "deeppink3"},
		{colorId: 162, hex: "#d70087", rgb: "rgb(215,0,135)", hsl: "hsl(322.325581395349%,100%,42%)", name: "deeppink3"},
		{colorId: 163, hex: "#d700af", rgb: "rgb(215,0,175)", hsl: "hsl(311.162790697674%,100%,42%)", name: "magenta3"},
		{colorId: 164, hex: "#d700d7", rgb: "rgb(215,0,215)", hsl: "hsl(300%,100%,42%)", name: "magenta3"},
		{colorId: 165, hex: "#d700ff", rgb: "rgb(215,0,255)", hsl: "hsl(290.588235294118%,100%,50%)", name: "magenta2"},
		{colorId: 166, hex: "#d75f00", rgb: "rgb(215,95,0)", hsl: "hsl(26.5116279069767%,100%,42%)", name: "darkorange3"},
		{colorId: 167, hex: "#d75f5f", rgb: "rgb(215,95,95)", hsl: "hsl(0%,60%,60%)", name: "indianred"},
		{colorId: 168, hex: "#d75f87", rgb: "rgb(215,95,135)", hsl: "hsl(340%,60%,60%)", name: "hotpink3"},
		{colorId: 169, hex: "#d75faf", rgb: "rgb(215,95,175)", hsl: "hsl(320%,60%,60%)", name: "hotpink2"},
		{colorId: 170, hex: "#d75fd7", rgb: "rgb(215,95,215)", hsl: "hsl(300%,60%,60%)", name: "orchid"},
		{colorId: 171, hex: "#d75fff", rgb: "rgb(215,95,255)", hsl: "hsl(285%,100%,68%)", name: "mediumorchid1"},
		{colorId: 172, hex: "#d78700", rgb: "rgb(215,135,0)", hsl: "hsl(37.6744186046512%,100%,42%)", name: "orange"},
		{colorId: 173, hex: "#d7875f", rgb: "rgb(215,135,95)", hsl: "hsl(20%,60%,60%)", name: "lightsalmon3"},
		{colorId: 174, hex: "#d78787", rgb: "rgb(215,135,135)", hsl: "hsl(0%,50%,68%)", name: "lightpink3"},
		{colorId: 175, hex: "#d787af", rgb: "rgb(215,135,175)", hsl: "hsl(330%,50%,68%)", name: "pink3"},
		{colorId: 176, hex: "#d787d7", rgb: "rgb(215,135,215)", hsl: "hsl(300%,50%,68%)", name: "plum3"},
		{colorId: 177, hex: "#d787ff", rgb: "rgb(215,135,255)", hsl: "hsl(280%,100%,76%)", name: "violet"},
		{colorId: 178, hex: "#d7af00", rgb: "rgb(215,175,0)", hsl: "hsl(48.8372093023256%,100%,42%)", name: "gold3"},
		{colorId: 179, hex: "#d7af5f", rgb: "rgb(215,175,95)", hsl: "hsl(40%,60%,60%)", name: "lightgoldenrod3"},
		{colorId: 180, hex: "#d7af87", rgb: "rgb(215,175,135)", hsl: "hsl(30%,50%,68%)", name: "tan"},
		{colorId: 181, hex: "#d7afaf", rgb: "rgb(215,175,175)", hsl: "hsl(0%,33%,76%)", name: "mistyrose3"},
		{colorId: 182, hex: "#d7afd7", rgb: "rgb(215,175,215)", hsl: "hsl(300%,33%,76%)", name: "thistle3"},
		{colorId: 183, hex: "#d7afff", rgb: "rgb(215,175,255)", hsl: "hsl(270%,100%,84%)", name: "plum2"},
		{colorId: 184, hex: "#d7d700", rgb: "rgb(215,215,0)", hsl: "hsl(60%,100%,42%)", name: "yellow3"},
		{colorId: 185, hex: "#d7d75f", rgb: "rgb(215,215,95)", hsl: "hsl(60%,60%,60%)", name: "khaki3"},
		{colorId: 186, hex: "#d7d787", rgb: "rgb(215,215,135)", hsl: "hsl(60%,50%,68%)", name: "lightgoldenrod2"},
		{colorId: 187, hex: "#d7d7af", rgb: "rgb(215,215,175)", hsl: "hsl(60%,33%,76%)", name: "lightyellow3"},
		{colorId: 188, hex: "#d7d7d7", rgb: "rgb(215,215,215)", hsl: "hsl(0%,0%,84%)", name: "grey84"},
		{colorId: 189, hex: "#d7d7ff", rgb: "rgb(215,215,255)", hsl: "hsl(240%,100%,92%)", name: "lightsteelblue1"},
		{colorId: 190, hex: "#d7ff00", rgb: "rgb(215,255,0)", hsl: "hsl(69.4117647058823%,100%,50%)", name: "yellow2"},
		{colorId: 191, hex: "#d7ff5f", rgb: "rgb(215,255,95)", hsl: "hsl(75%,100%,68%)", name: "darkolivegreen1"},
		{colorId: 192, hex: "#d7ff87", rgb: "rgb(215,255,135)", hsl: "hsl(80%,100%,76%)", name: "darkolivegreen1"},
		{colorId: 193, hex: "#d7ffaf", rgb: "rgb(215,255,175)", hsl: "hsl(90%,100%,84%)", name: "darkseagreen1"},
		{colorId: 194, hex: "#d7ffd7", rgb: "rgb(215,255,215)", hsl: "hsl(120%,100%,92%)", name: "honeydew2"},
		{colorId: 195, hex: "#d7ffff", rgb: "rgb(215,255,255)", hsl: "hsl(180%,100%,92%)", name: "lightcyan1"},
		{colorId: 196, hex: "#ff0000", rgb: "rgb(255,0,0)", hsl: "hsl(0%,100%,50%)", name: "red1"},
		{colorId: 197, hex: "#ff005f", rgb: "rgb(255,0,95)", hsl: "hsl(337.647058823529%,100%,50%)", name: "deeppink2"},
		{colorId: 198, hex: "#ff0087", rgb: "rgb(255,0,135)", hsl: "hsl(328.235294117647%,100%,50%)", name: "deeppink1"},
		{colorId: 199, hex: "#ff00af", rgb: "rgb(255,0,175)", hsl: "hsl(318.823529411765%,100%,50%)", name: "deeppink1"},
		{colorId: 200, hex: "#ff00d7", rgb: "rgb(255,0,215)", hsl: "hsl(309.411764705882%,100%,50%)", name: "magenta2"},
		{colorId: 201, hex: "#ff00ff", rgb: "rgb(255,0,255)", hsl: "hsl(300%,100%,50%)", name: "magenta1"},
		{colorId: 202, hex: "#ff5f00", rgb: "rgb(255,95,0)", hsl: "hsl(22.3529411764706%,100%,50%)", name: "orangered1"},
		{colorId: 203, hex: "#ff5f5f", rgb: "rgb(255,95,95)", hsl: "hsl(0%,100%,68%)", name: "indianred1"},
		{colorId: 204, hex: "#ff5f87", rgb: "rgb(255,95,135)", hsl: "hsl(345%,100%,68%)", name: "indianred1"},
		{colorId: 205, hex: "#ff5faf", rgb: "rgb(255,95,175)", hsl: "hsl(330%,100%,68%)", name: "hotpink"},
		{colorId: 206, hex: "#ff5fd7", rgb: "rgb(255,95,215)", hsl: "hsl(315%,100%,68%)", name: "hotpink"},
		{colorId: 207, hex: "#ff5fff", rgb: "rgb(255,95,255)", hsl: "hsl(300%,100%,68%)", name: "mediumorchid1"},
		{colorId: 208, hex: "#ff8700", rgb: "rgb(255,135,0)", hsl: "hsl(31.7647058823529%,100%,50%)", name: "darkorange"},
		{colorId: 209, hex: "#ff875f", rgb: "rgb(255,135,95)", hsl: "hsl(15%,100%,68%)", name: "salmon1"},
		{colorId: 210, hex: "#ff8787", rgb: "rgb(255,135,135)", hsl: "hsl(0%,100%,76%)", name: "lightcoral"},
		{colorId: 211, hex: "#ff87af", rgb: "rgb(255,135,175)", hsl: "hsl(340%,100%,76%)", name: "palevioletred1"},
		{colorId: 212, hex: "#ff87d7", rgb: "rgb(255,135,215)", hsl: "hsl(320%,100%,76%)", name: "orchid2"},
		{colorId: 213, hex: "#ff87ff", rgb: "rgb(255,135,255)", hsl: "hsl(300%,100%,76%)", name: "orchid1"},
		{colorId: 214, hex: "#ffaf00", rgb: "rgb(255,175,0)", hsl: "hsl(41.1764705882353%,100%,50%)", name: "orange1"},
		{colorId: 215, hex: "#ffaf5f", rgb: "rgb(255,175,95)", hsl: "hsl(30%,100%,68%)", name: "sandybrown"},
		{colorId: 216, hex: "#ffaf87", rgb: "rgb(255,175,135)", hsl: "hsl(20%,100%,76%)", name: "lightsalmon1"},
		{colorId: 217, hex: "#ffafaf", rgb: "rgb(255,175,175)", hsl: "hsl(0%,100%,84%)", name: "lightpink1"},
		{colorId: 218, hex: "#ffafd7", rgb: "rgb(255,175,215)", hsl: "hsl(330%,100%,84%)", name: "pink1"},
		{colorId: 219, hex: "#ffafff", rgb: "rgb(255,175,255)", hsl: "hsl(300%,100%,84%)", name: "plum1"},
		{colorId: 220, hex: "#ffd700", rgb: "rgb(255,215,0)", hsl: "hsl(50.5882352941176%,100%,50%)", name: "gold1"},
		{colorId: 221, hex: "#ffd75f", rgb: "rgb(255,215,95)", hsl: "hsl(45%,100%,68%)", name: "lightgoldenrod2"},
		{colorId: 222, hex: "#ffd787", rgb: "rgb(255,215,135)", hsl: "hsl(40%,100%,76%)", name: "lightgoldenrod2"},
		{colorId: 223, hex: "#ffd7af", rgb: "rgb(255,215,175)", hsl: "hsl(30%,100%,84%)", name: "navajowhite1"},
		{colorId: 224, hex: "#ffd7d7", rgb: "rgb(255,215,215)", hsl: "hsl(0%,100%,92%)", name: "mistyrose1"},
		{colorId: 225, hex: "#ffd7ff", rgb: "rgb(255,215,255)", hsl: "hsl(300%,100%,92%)", name: "thistle1"},
		{colorId: 226, hex: "#ffff00", rgb: "rgb(255,255,0)", hsl: "hsl(60%,100%,50%)", name: "yellow1"},
		{colorId: 227, hex: "#ffff5f", rgb: "rgb(255,255,95)", hsl: "hsl(60%,100%,68%)", name: "lightgoldenrod1"},
		{colorId: 228, hex: "#ffff87", rgb: "rgb(255,255,135)", hsl: "hsl(60%,100%,76%)", name: "khaki1"},
		{colorId: 229, hex: "#ffffaf", rgb: "rgb(255,255,175)", hsl: "hsl(60%,100%,84%)", name: "wheat1"},
		{colorId: 230, hex: "#ffffd7", rgb: "rgb(255,255,215)", hsl: "hsl(60%,100%,92%)", name: "cornsilk1"},
		{colorId: 231, hex: "#ffffff", rgb: "rgb(255,255,255)", hsl: "hsl(0%,0%,100%)", name: "grey100"},
		{colorId: 232, hex: "#080808", rgb: "rgb(8,8,8)", hsl: "hsl(0%,0%,3%)", name: "grey3"},
		{colorId: 233, hex: "#121212", rgb: "rgb(18,18,18)", hsl: "hsl(0%,0%,7%)", name: "grey7"},
		{colorId: 234, hex: "#1c1c1c", rgb: "rgb(28,28,28)", hsl: "hsl(0%,0%,10%)", name: "grey11"},
		{colorId: 235, hex: "#262626", rgb: "rgb(38,38,38)", hsl: "hsl(0%,0%,14%)", name: "grey15"},
		{colorId: 236, hex: "#303030", rgb: "rgb(48,48,48)", hsl: "hsl(0%,0%,18%)", name: "grey19"},
		{colorId: 237, hex: "#3a3a3a", rgb: "rgb(58,58,58)", hsl: "hsl(0%,0%,22%)", name: "grey23"},
		{colorId: 238, hex: "#444444", rgb: "rgb(68,68,68)", hsl: "hsl(0%,0%,26%)", name: "grey27"},
		{colorId: 239, hex: "#4e4e4e", rgb: "rgb(78,78,78)", hsl: "hsl(0%,0%,30%)", name: "grey30"},
		{colorId: 240, hex: "#585858", rgb: "rgb(88,88,88)", hsl: "hsl(0%,0%,34%)", name: "grey35"},
		{colorId: 241, hex: "#626262", rgb: "rgb(98,98,98)", hsl: "hsl(0%,0%,37%)", name: "grey39"},
		{colorId: 242, hex: "#6c6c6c", rgb: "rgb(108,108,108)", hsl: "hsl(0%,0%,40%)", name: "grey42"},
		{colorId: 243, hex: "#767676", rgb: "rgb(118,118,118)", hsl: "hsl(0%,0%,46%)", name: "grey46"},
		{colorId: 244, hex: "#808080", rgb: "rgb(128,128,128)", hsl: "hsl(0%,0%,50%)", name: "grey50"},
		{colorId: 245, hex: "#8a8a8a", rgb: "rgb(138,138,138)", hsl: "hsl(0%,0%,54%)", name: "grey54"},
		{colorId: 246, hex: "#949494", rgb: "rgb(148,148,148)", hsl: "hsl(0%,0%,58%)", name: "grey58"},
		{colorId: 247, hex: "#9e9e9e", rgb: "rgb(158,158,158)", hsl: "hsl(0%,0%,61%)", name: "grey62"},
		{colorId: 248, hex: "#a8a8a8", rgb: "rgb(168,168,168)", hsl: "hsl(0%,0%,65%)", name: "grey66"},
		{colorId: 249, hex: "#b2b2b2", rgb: "rgb(178,178,178)", hsl: "hsl(0%,0%,69%)", name: "grey70"},
		{colorId: 250, hex: "#bcbcbc", rgb: "rgb(188,188,188)", hsl: "hsl(0%,0%,73%)", name: "grey74"},
		{colorId: 251, hex: "#c6c6c6", rgb: "rgb(198,198,198)", hsl: "hsl(0%,0%,77%)", name: "grey78"},
		{colorId: 252, hex: "#d0d0d0", rgb: "rgb(208,208,208)", hsl: "hsl(0%,0%,81%)", name: "grey82"},
		{colorId: 253, hex: "#dadada", rgb: "rgb(218,218,218)", hsl: "hsl(0%,0%,85%)", name: "grey85"},
		{colorId: 254, hex: "#e4e4e4", rgb: "rgb(228,228,228)", hsl: "hsl(0%,0%,89%)", name: "grey89"},
		{colorId: 255, hex: "#eeeeee", rgb: "rgb(238,238,238)", hsl: "hsl(0%,0%,93%)", name: "grey93"},
	}

	for _, c := range colors {
		if strings.HasPrefix(s, "#") && c.hex == s {
			return c.colorId
		} else if strings.HasPrefix(s, "rgb") && c.rgb == s {
			return c.colorId
		} else if strings.HasPrefix(s, "hsl") && c.hsl == s {
			return c.colorId
		} else if c.name == s {
			return c.colorId
		}
	}
	return 15
}