package services

//// SendMsg 发送信息
//func SendMsg(addr, order int, c string) {
//	global.Port.SendData(addr, order, 0, []byte(c))
//	global.Logger.Info("Lora send msg",
//		zap.String("time", time.Now().Format("2006-01-02 15:04:05")),
//		zap.Namespace("LoraMsg"),
//		zap.Int("address", addr),
//		zap.String("msg", c),
//	)
//}
//
//func ReceiveImg(path string) {
//	sum := -1
//	img := make([]byte, 0)
//	begin := time.Now()
//	for {
//		data, flag, err := global.Port.ReceiveData()
//		fmt.Println("receive: ", flag)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		if flag == 1 {
//			sum = int(data[0])<<24 + int(data[1])<<16 + int(data[2])<<8 + int(data[3])
//			break
//		}
//	}
//
//	for {
//		data, flag, err := global.Port.ReceiveData()
//		fmt.Println("receive: ", flag)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		img = append(img, data...)
//		if flag == sum {
//			break
//		}
//	}
//	os.WriteFile(path, img, 0777)
//	fmt.Println(time.Since(begin))
//
//}
