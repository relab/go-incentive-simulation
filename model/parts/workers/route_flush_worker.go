package workers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-incentive-simulation/model/parts/types"
	"os"
	"sync"
)

func RouteFlushWorker(routeChan chan types.RouteData, wg *sync.WaitGroup) {
	defer wg.Done()
	//var message *protoGenerated.RouteData
	var routeData types.RouteData
	var bytes []byte
	filePath := "./results/routes.txt"

	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("No need to remove file with path: ", filePath)
	}

	actualFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer func(actualFile *os.File) {
		err1 := actualFile.Close()
		if err1 != nil {
			fmt.Println("Couldn't close the file with filepath: ", filePath)
		}
	}(actualFile)

	writer := bufio.NewWriter(actualFile) // default writer size is 4096 bytes
	//writer = bufio.NewWriterSize(writer, 1048576) // 1MiB
	defer func(writer *bufio.Writer) {
		err1 := writer.Flush()
		if err1 != nil {
			fmt.Println("Couldn't flush the remaining buffer in the writer for states")
		}
	}(writer)

	for routeData = range routeChan {

		bytes, err = json.Marshal(routeData)
		_, err = writer.Write(bytes)
		if err != nil {
			panic(err)
		}
		_, err = writer.WriteString("\n")
		if err != nil {
			panic(err)
		}

		// TODO: uncomment below to use messagePack
		//bytes, err = msgpack.Marshal(routeData)
		//if err != nil {
		//	panic(err)
		//}

		// TODO: uncomment below to write to binary
		//message = &protoGenerated.RouteData{
		//	TimeStep: routeData.TimeStep,
		//	RequestResult:    &protoGenerated.RequestResult{},
		//}
		//for _, nodeId := range routeData.RequestResult {
		//	message.RequestResult.Waypoints = append(message.RequestResult.Waypoints, int32(nodeId))
		//}
		//
		//bytes, err = proto.Marshal(message)
		//if err != nil {
		//	panic(err)
		//}
	}
}

//func routeListConvertAndDumpToFile(routes []types.RequestResult, curTimeStep int, actualFile *os.File) error {

// message := &protoGenerated.RouteData{
// 	TimeStep: int32(curTimeStep),
// 	Routes:   make([]*protoGenerated.RequestResult, len(routes)),
// }
// for i, route := range routes {
// 	var routeList []int32
// 	for _, node := range route {
// 		routeList = append(routeList, int32(node))
// 	}
// 	message.Routes[i] = &protoGenerated.RequestResult{
// 		Waypoints: routeList,
// 	}
// }
// data1, err := proto.Marshal(message)
//data := RouteData{curTimeStep, routes}
//file, _ := json.Marshal(data)
////actualFile, err := os.OpenFile("routes.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
////if err != nil {
////	return err
////}
//w := bufio.NewWriter(actualFile)
//_, err := w.Write(file)
//if err != nil {
//	panic(err)
//}
//err = w.Flush()
//if err != nil {
//	panic(err)
//}
//return nil
//}

//func routeListAndFlush(state *types.State, route types.RequestResult, curTimeStep int, actualFile *os.File) []types.RequestResult {
//	state.RouteLists[curTimeStep%10000] = route
//	if (curTimeStep+5000)%10000 == 0 {
//		routeListConvertAndDumpToFile(state.RouteLists, curTimeStep, actualFile)
//		state.RouteLists = make([]types.RequestResult, 10000)
//	}
//	return state.RouteLists
//}
