package convnetlib

func StartConsole() {

	//创建本地UDP服务
	client.UdpServerPort, client.g_udpserver = SatrtUDPServer(8080, 10)
	//尝试打开UPNP
	UdpServerUpnpSet(client.UdpServerPort)
	//创建本地HTTP-API服务
	StartHttpServer(8081, 10)
}
