package gsheet

import (
	"google.golang.org/api/sheets/v4"
)

func createBasicChartSeries(values [][]interface{}) []*sheets.BasicChartSeries {
	var BasicChartSeries []*sheets.BasicChartSeries
	for i := range values {
		if i == 0 {
			continue
		}
		BasicChartSeries = append(BasicChartSeries, &sheets.BasicChartSeries{
			Series: &sheets.ChartData{
				SourceRange: &sheets.ChartSourceRange{
					Sources: []*sheets.GridRange{
						{ //A2:O2
							SheetId:          1023,
							StartRowIndex:    int64(i),
							EndRowIndex:      int64(i + 1),
							StartColumnIndex: 0,
							EndColumnIndex:   int64(len(values[i])),
						},
					},
				},
			},
		})
	}
	return BasicChartSeries
}

func createHistogramSeries(values [][]interface{}) []*sheets.HistogramSeries {
	var HistogramSeries []*sheets.HistogramSeries
	for i := range values {
		if i == 0 {
			continue
		}
		HistogramSeries = append(HistogramSeries, &sheets.HistogramSeries{
			Data: &sheets.ChartData{
				SourceRange: &sheets.ChartSourceRange{
					Sources: []*sheets.GridRange{
						{ //A2:O2
							SheetId:          1023,
							StartRowIndex:    int64(i),
							EndRowIndex:      int64(i + 1),
							StartColumnIndex: 0,
							EndColumnIndex:   int64(len(values[i])),
						},
					},
				},
			},
		})
	}
	return HistogramSeries
}

func createBasicChartSpec(legendPosition, chartype, stackedType string, sourceSheetId int64, values [][]interface{}, basicChartSeries []*sheets.BasicChartSeries) *sheets.BasicChartSpec {
	var BasicChartSpec *sheets.BasicChartSpec
	BasicChartSpec = &sheets.BasicChartSpec{
		HeaderCount:    1,
		Series:         basicChartSeries,
		LegendPosition: "BOTTOM_LEGEND",
		ChartType:      "COLUMN",
		StackedType:    "STACKED",
		Domains: []*sheets.BasicChartDomain{
			{
				Domain: &sheets.ChartData{
					SourceRange: &sheets.ChartSourceRange{
						Sources: []*sheets.GridRange{
							{
								SheetId:          1023,
								StartRowIndex:    0,
								EndRowIndex:      1,
								StartColumnIndex: 0,
								EndColumnIndex:   int64(len(values[0])),
							},
						},
					},
				},
			},
		},
	}
	return BasicChartSpec
}
