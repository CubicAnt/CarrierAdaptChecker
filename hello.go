package main

import (
	"CarrierAdaptChecker/util"
	"flag"
	"log"
	"os"
	"strings"
)

var (
	workSpace string
	path string
)

func parseFlag() {
	flag.StringVar(&workSpace, "d", "", "workspace path")
	flag.StringVar(&path, "p", "checklist.txt", "the checklist file for check, format like Project: XXX\nChange-Id: XXX\n")
	flag.Parse()
	if len(workSpace) == 0 {
		log.Fatal("please input workspace path")
	}
}

func main() {
	parseFlag()
	changeMap := getChangeMapByProject()
	statisticsCnt(changeMap)
}

func statisticsCnt(changeMap map[string][]string) {
	allCnt := 0
	failCnt := 0
	successCnt := 0
	for project, changeList := range changeMap {
		allCnt += len(changeList)
		if util.RepoSyncProject(workSpace, project) {
			projectWorkSpace := workSpace + string(os.PathSeparator) + project
			util.GitPullRebase(projectWorkSpace)
			logList := util.GitLog(projectWorkSpace)
			for _, change := range changeList {
				if !strings.Contains(logList, change) {
					failCnt++
					log.Println("Project: ", project, " Change-Id: ", change, " check fail!!!")
				} else {
					successCnt++
				}
			}
		}
	}

	if failCnt == 0 && successCnt == allCnt {
		log.Println("Congratulations! All Pass(", successCnt, "/", allCnt, ")")
	} else {
		log.Println("Not OK! pass statistics: failCnt: ", failCnt, " successCnt: ", successCnt, " allCnt: ", allCnt)
	}
}

func getChangeMapByProject() map[string][]string {
	kwProject := "Project: "
	kwChangeId := "Change-Id: "
	changeMap := make(map[string][]string)
	tempProject := ""
	util.ReadLine(path, func(text string) {
		switch {
		case strings.Contains(text, kwProject):
			tempProject = strings.TrimSpace(text[len(kwProject):])
		case strings.Contains(text, kwChangeId):
			if len(tempProject) > 0 {
				changeMap[tempProject] = append(changeMap[tempProject], strings.TrimSpace(text[len(kwChangeId):]))
			}
		}
	})
	return changeMap
}
