package mongo

import (
	"gopkg.in/mgo.v2"
)

func NewSession(mongoUrl string) Session {
	mgoSession, err := mgo.Dial(mongoUrl)
	if err != nil {
		panic(err)
	}
	return MongoSession{mgoSession}
}

func NewMockSession(mockSessionResults MockSessionResults, mockDatabaseResults map[string]MockDatabaseResults) Session {
	return MockSession{mockSessionResults, mockDatabaseResults}
}

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
		"event_type":                                                              "LoadBalancerSample",
		"provider":                                                                PROVIDER,
		"mongo.server.host":                                                       serverStatus.Host,
		"mongo.server.version":                                                    serverStatus.Version,
		"mongo.server.pid":                                                        serverStatus.Pid,
		"mongo.server.uptime":                                                     serverStatus.Uptime,
		"mongo.server.uptimeMillis":                                               serverStatus.UptimeMillis,
		"mongo.server.uptimeEstimate":                                             serverStatus.UptimeEstimate,
		"mongo.server.asserts.msg":                                                serverStatus.Asserts.Msg,
		"mongo.server.asserts.regular":                                            serverStatus.Asserts.Regular,
		"mongo.server.asserts.rollovers":                                          serverStatus.Asserts.Rollovers,
		"mongo.server.asserts.user":                                               serverStatus.Asserts.User,
		"mongo.server.asserts.warning":                                            serverStatus.Asserts.Warning,
		"mongo.backgroundFlushing.averageMS":                                      serverStatus.BackgroundFlushing.AverageMS,
		"mongo.backgroundFlushing.flushes":                                        serverStatus.BackgroundFlushing.Flushes,
		"mongo.backgroundFlushing.lastMS.":                                        serverStatus.BackgroundFlushing.LastMS,
		"mongo.backgroundFlushing.totalMS":                                        serverStatus.BackgroundFlushing.TotalMS,
		"mongo.connections.available":                                             serverStatus.Connections.Available,
		"mongo.connections.current":                                               serverStatus.Connections.Current,
		"mongo.connections.totalCreated":                                          serverStatus.Connections.TotalCreated,
		"mongo.dur.commits":                                                       serverStatus.Dur.Commits,
		"mongo.dur.compression":                                                   serverStatus.Dur.Compression,
		"mongo.dur.earlyCommits":                                                  serverStatus.Dur.EarlyCommits,
		"mongo.dur.journalMB":                                                     serverStatus.Dur.JournaledMB,
		"mongo.dur.writeToDataFilesMb":                                            serverStatus.Dur.WriteToDataFilesMB,
		"mongo.dur.commitsInWriteLock":                                            serverStatus.Dur.CommitsInWriteLock,
		"mongo.dur.timems.dt":                                                     serverStatus.Dur.TimeMS.DT,
		"mongo.dur.timems.prepLogBuffer":                                          serverStatus.Dur.TimeMS.PrepLogBuffer,
		"mongo.dur.timems.writeToJournal":                                         serverStatus.Dur.TimeMS.WriteToJournal,
		"mongo.dur.timems.writeToDataFiles":                                       serverStatus.Dur.TimeMS.WriteToDataFiles,
		"mongo.dur.timems.remapPrivateView":                                       serverStatus.Dur.TimeMS.RemapPrivateView,
		"mongo.dur.timems.commits":                                                serverStatus.Dur.TimeMS.Commits,
		"mongo.dur.timems.commitsInWriteLock":                                     serverStatus.Dur.TimeMS.CommitsInWriteLock,
		"mongo.extraInfo.pageFaults":                                              serverStatus.ExtraInfo.PageFaults,
		"mongo.globalLock.totalTime":                                              serverStatus.GlobalLock.TotalTime,
		"mongo.globalLock.currentQueue.totalTime":                                 serverStatus.GlobalLock.CurrentQueue.Total,
		"mongo.globalLock.currentQueue.readers":                                   serverStatus.GlobalLock.CurrentQueue.Readers,
		"mongo.globalLock.currentQueue.writers":                                   serverStatus.GlobalLock.CurrentQueue.Writers,
		"mongo.globalLock.activeClients.totalTime":                                serverStatus.GlobalLock.ActiveClients.Total,
		"mongo.globalLock.activeClients.readers":                                  serverStatus.GlobalLock.ActiveClients.Readers,
		"mongo.globalLock.activeClients.writers":                                  serverStatus.GlobalLock.ActiveClients.Writers,
		"mongo.network.bytesIn":                                                   serverStatus.Network.BytesIn,
		"mongo.network.bytesOut":                                                  serverStatus.Network.BytesOut,
		"mongo.network.requests":                                                  serverStatus.Network.NumRequests,
		"mongo.opcounters.insert":                                                 serverStatus.OpCounters.Insert,
		"mongo.opcounters.query":                                                  serverStatus.OpCounters.Query,
		"mongo.opcounters.update":                                                 serverStatus.OpCounters.Update,
		"mongo.opcounters.delete":                                                 serverStatus.OpCounters.Delete,
		"mongo.opcounters.getmore":                                                serverStatus.OpCounters.Getmore,
		"mongo.opcounters.command":                                                serverStatus.OpCounters.Command,
		"mongo.storageEngine.name":                                                serverStatus.StorageEngine.Name,
		"mongo.storageEngine.supportsCommittedReads":                              serverStatus.StorageEngine.SupportsCommittedReads,
		"mongo.storageEngine.persistent":                                          serverStatus.StorageEngine.Persistent,
		"mongo.wiredTiger.cache.bytesCurrentlyInCache":                            serverStatus.WiredTiger.Cache.BytesCurrentlyInCache,
		"mongo.wiredTiger.cache.failedEvictionPagesExceedingTheInMemoryMaximumps": serverStatus.WiredTiger.Cache.FailedEvictionPagesExceedingTheInMemoryMaximumps,
		"mongo.wiredTiger.cache.inMemoryPageSplits":                               serverStatus.WiredTiger.Cache.InMemoryPageSplits,
		"mongo.wiredTiger.cache.maximumBytesConfigured":                           serverStatus.WiredTiger.Cache.MaximumBytesConfigured,
		"mongo.wiredTiger.cache.maximumPageSizeAtEviction":                        serverStatus.WiredTiger.Cache.MaximumPageSizeAtEviction,
		"mongo.wiredTiger.cache.modifiedPagesEvicted":                             serverStatus.WiredTiger.Cache.ModifiedPagesEvicted,
		"mongo.wiredTiger.cache.pagesCurrentlyHeldInTheCache":                     serverStatus.WiredTiger.Cache.PagesCurrentlyHeldInTheCache,
		"mongo.wiredTiger.cache.pagesEvictedByApplicationThreads":                 serverStatus.WiredTiger.Cache.PagesEvictedByApplicationThreads,
		"mongo.wiredTiger.cache.pagesEvictedBecauseTheyExeceededTheInMemoryMax":   serverStatus.WiredTiger.Cache.PagesEvictedBecauseTheyExeceededTheInMemoryMax,
		"mongo.wiredTiger.cache.trackedDirtyBytesInTheCache":                      serverStatus.WiredTiger.Cache.TrackedDirtyBytesInTheCache,
		"mongo.wiredTiger.cache.unmodifiedPagesEvicted":                           serverStatus.WiredTiger.Cache.UnmodifiedPagesEvicted,
		"mongo.wiredTiger.concurrentTransactions.write.out":                       serverStatus.WiredTiger.ConcurrentTransations.Write.Out,
		"mongo.wiredTiger.concurrentTransactions.write.available":                 serverStatus.WiredTiger.ConcurrentTransations.Write.Available,
		"mongo.wiredTiger.concurrentTransactions.write.totalTickets":              serverStatus.WiredTiger.ConcurrentTransations.Write.TotalTickets,
		"mongo.wiredTiger.concurrentTransactions.read.out":                        serverStatus.WiredTiger.ConcurrentTransations.Read.Out,
		"mongo.wiredTiger.concurrentTransactions.read.available":                  serverStatus.WiredTiger.ConcurrentTransations.Read.Available,
		"mongo.wiredTiger.concurrentTransactions.read.totalTickets":               serverStatus.WiredTiger.ConcurrentTransations.Read.TotalTickets,
		"mongo.mem.bits":                                                          serverStatus.Mem.Bits,
		"mongo.mem.resident":                                                      serverStatus.Mem.Resident,
		"mongo.mem.virtual":                                                       serverStatus.Mem.Virtual,
		"mongo.mem.supported":                                                     serverStatus.Mem.Supported,
		"mongo.mem.mapped":                                                        serverStatus.Mem.Mapped,
		"mongo.mem.mappedWithJournal":                                             serverStatus.Mem.MappedWithJournal,
		"mongo.metrics.cursor.timedOut":                                           serverStatus.Metrics.Cursor.TimedOut,
		"mongo.metrics.cursor.open.noTimeout":                                     serverStatus.Metrics.Cursor.Open.NoTimeout,
		"mongo.metrics.cursor.open.pinned":                                        serverStatus.Metrics.Cursor.Open.Pinned,
		"mongo.metrics.cursor.open.Total":                                         serverStatus.Metrics.Cursor.Open.Total,
		"mongo.metrics.document.deleted":                                          serverStatus.Metrics.Document.Deleted,
		"mongo.metrics.document.inserted":                                         serverStatus.Metrics.Document.Inserted,
		"mongo.metrics.document.updated":                                          serverStatus.Metrics.Document.Updated,
		"mongo.metrics.document.returned":                                         serverStatus.Metrics.Document.Returned,
		"mongo.metrics.getLastError.wtimeouts":                                    serverStatus.Metrics.GetLastError.Wtimeouts,
		"mongo.metrics.getLastError.wtime.num":                                    serverStatus.Metrics.GetLastError.Wtime.Num,
		"mongo.metrics.getLastError.wtime.TotalMillis":                            serverStatus.Metrics.GetLastError.Wtime.TotalMillis,
		"mongo.metrics.operation.fastmod":                                         serverStatus.Metrics.Operation.Fastmod,
		"mongo.metrics.operation.idhack":                                          serverStatus.Metrics.Operation.Idhack,
		"mongo.metrics.operation.scanAndOrder":                                    serverStatus.Metrics.Operation.ScanAndOrder,
		"mongo.metrics.operation.writeConflicts":                                  serverStatus.Metrics.Operation.WriteConflicts,
		"mongo.metrics.queryExecutor.scanned":                                     serverStatus.Metrics.QueryExecutor.Scanned,
		"mongo.metrics.queryExecutor.scannedObjects":                              serverStatus.Metrics.QueryExecutor.ScannedObjects,
	}
}
