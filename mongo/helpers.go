package mongo

func formatDBStatsStructToMap(dbStats dbStats) (dbInfoMap map[string]interface{}) {
	return map[string]interface{}{
		"event_type":           "LoadBalancerSample",
		"provider":             PROVIDER,
		"mongo.db.name":        dbStats.DB,
		"mongo.db.collections": dbStats.Collections,
		"mongo.db.Objects":     dbStats.Objects,
		"mongo.db.AvgObjSize":  dbStats.AvgObjSize,
		"mongo.db.DataSize":    dbStats.DataSize,
		"mongo.db.StorageSize": dbStats.StorageSize,
		"mongo.db.NumExtents":  dbStats.NumExtents,
		"mongo.db.Indexes":     dbStats.NumExtents,
		"mongo.db.IndexSize":   dbStats.IndexSize,
	}
}

func formatServerStatsStructToMap(serverStatus serverStatus) (dbInfoMap map[string]interface{}) {
	return map[string]interface{}{
		"event_type":                     "LoadBalancerSample",
		"provider":                       PROVIDER,
		"mongo.server.host":              serverStatus.Host,
		"mongo.server.version":           serverStatus.Version,
		"mongo.server.pid":               serverStatus.Pid,
		"mongo.server.uptime":            serverStatus.Uptime,
		"mongo.server.uptimeMillis":      serverStatus.UptimeMillis,
		"mongo.server.uptimeEstimate":    serverStatus.UptimeEstimate,
		"mongo.server.asserts.msg":       serverStatus.Asserts.Msg,
		"mongo.server.asserts.regular":   serverStatus.Asserts.Regular,
		"mongo.server.asserts.rollovers": serverStatus.Asserts.Rollovers,
		"mongo.server.asserts.user":      serverStatus.Asserts.User,
		"mongo.server.asserts.warning":   serverStatus.Asserts.Warning,
	}
}
