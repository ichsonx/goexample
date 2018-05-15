package main

/*
	生成注册登记薄履约配额导入的sql脚本
	1.准备与政府给予数据对比完全的数据文件，并且按照列序：account_id，企业名称，当年履约配额数量，放到Sheet1中
*/
import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
	"os"
	"io"
	"time"
)

var(
	xlsxpath= "./examplefiles/final.xlsx"
	xlsxarr [][]string
	sqlcontent = ""
	sqlpath = "./examplefiles/final.sql"
)

func main() {
	//判断履约配额导入的xlsx文件是否存在，存在则声场导入的sql文件
	if _, err := os.Stat(xlsxpath); err == nil{
		fmt.Printf("开始生成履约配额导入文件sql\n")
		initxlsx()
		header := "prompt PL/SQL Developer import file \nprompt Created on 2015年5月21日 by Jenny \nset feedback off \nset define off\nprompt Loading ORG_CARBON_EMISSIONS_INFO...\n"
		header2 := "insert into ORG_CARBON_EMISSIONS_INFO "
		header3 := "(ID, ACCOUNT_ID1, ACCOUNT_ID2, ACCOUNT_NAME, ALLOWANCE_TYPE, EMISSION_DATA, TERMYEARS, TERMYEARE, CREATE_USER, CREATE_TIME, UPDATE_USER, UPDATE_TIME) values ("
		header4 := header2 + header3
		beginnum := 2800
		termyear1 := "to_date('01-01-2017', 'dd-mm-yyyy')"	//履约开始年份，一般2018年进行2017年的履约
		termyear2 := "to_date('31-12-2017', 'dd-mm-yyyy')"	//履约结束年份
		creattime := "to_timestamp('21-05-2018 19:20:14.000000', 'dd-mm-yyyy hh24:mi:ss.ff')"	//这个操作创建和更新时间
		tail := fmt.Sprintf("commit;\nprompt %d records loaded\nset feedback on\nset define on\nprompt Done.", len(xlsxarr))

		file, err := os.OpenFile(sqlpath, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil{
			fmt.Printf("打开sqlpath失败：%v \n" , err)
		}
		_, _ = io.WriteString(file, header)
		for _, row := range xlsxarr{
			sqlcontent += fmt.Sprintf("%s %d, '%s', '%s', '%s', '11', %s, %s, %s, 'sys', %s, 'sys', %s);\n", header4, beginnum, row[0], row[0], row[1], row[2], termyear1, termyear2, creattime, creattime)
			beginnum++
		}
		sqlcontent += tail
		_, _ = io.WriteString(file, sqlcontent)
		file.Sync()
		fmt.Printf("生成履约配额导入sql文件完成\n")
	}
	time.Sleep(time.Minute * 10)
}

func initxlsx()  {
	xlsx, err := excelize.OpenFile(xlsxpath)
	if err != nil{
		fmt.Print(err)
	}
	xlsxarr = xlsx.GetRows("Sheet1")
}
