package common

import "testing"

func TestAll(t *testing.T) {
	t.Run("TestIsDate", TestIsDate)
}

func TestMain(m *testing.M) {
	m.Run()
}

type testIsDateData struct {
	date string
	result bool
}

func TestIsDate(t *testing.T) {
	var result bool

	var data []testIsDateData

	data = append(data, testIsDateData{
		date: "2021-05-06",
		result:    true,
	})

	data = append(data, testIsDateData{
		date: "2021-5-06",
		result:    true,
	})

	data = append(data, testIsDateData{
		date: "2021-05-6",
		result:    true,
	})

	data = append(data, testIsDateData{
		date: "2021/05/06",
		result:    true,
	})

	data = append(data, testIsDateData{
		date: "2021/5/06",
		result:    true,
	})

	data = append(data, testIsDateData{
		date: "2021/05/6",
		result:    true,
	})

	data = append(data, testIsDateData{
		date: "2021.05.6",
		result:    false,
	})

	data = append(data, testIsDateData{
		date: "1995[05]6",
		result:    false,
	})

	for _, d := range(data) {
		result = IsDate(d.date)
		if result != d.result {
			t.Errorf(`"%s"应该为 true, 测试结果为: %t`, d.date, result)
		}
	}
}

type testCompareTime struct {
	date1 string
	symbol string
	date2 string
	result bool
}

func TestCompareTime(t *testing.T) {
	var data []testCompareTime

	data = append(data, testCompareTime{
		date1:  "2021-04-05",
		symbol: ">",
		date2:  "2021-04-03",
		result: true,
	})

	data = append(data, testCompareTime{
		date1:  "2021-04-05",
		symbol: "<",
		date2:  "2021-04-06",
		result: true,
	})

	data = append(data, testCompareTime{
		date1:  "2021-04-05",
		symbol: "=",
		date2:  "2021-04-05",
		result: true,
	})

	data = append(data, testCompareTime{
		date1:  "2021-04-05",
		symbol: ">",
		date2:  "2021-04-06",
		result: false,
	})

	data = append(data, testCompareTime{
		date1:  "2021-04-06",
		symbol: "<",
		date2:  "2021-04-03",
		result: false,
	})

	for _, d := range(data) {
		result, err := CompareTime(d.date1, d.symbol, d.date2)
		if err != nil {
			t.Error(err)
		}

		if result != d.result {
			t.Errorf(`"%s %s %s" 结果应为 %t, 测试结果为: %t`, d.date1, d.symbol, d.date2, d.result, result)
		}
	}
}