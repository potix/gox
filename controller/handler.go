package controller

/*
   /gox POST gox終了する action=stop
   /gox/config GET gox自体の設定取得
   /gox/config PUT gox自体の設定変更
   /pulugins GET availableなプラグイン一覧 type=available
   /pulugins GET enableなプラグイン一覧 type=enable
   /pulugins GET ロードされてるなプラグイン一覧  type=loades
   /pulugins POST プラグインをダウンロードする action=download
   /pulugins POST プラグインを有効にする(symlink作る) action=load
   /pulugins POST プラグイン全体をロードする action=load
   /pulugins POST プラグイン全体をアンロードする action=unload
   /puligins/XXXX DELETE プラグイン単体をアンロードする
   /puligins/XXXX POST プラグインの操作系 action = up(new + init + start + basic command), action= down(stop + fini + free)
   /puligins/XXXX/help GET プラグインのヘルプ  command help
   /puligins/XXXX/commands GET プラグインのコマンド情報 command getCommands
   /puligins/XXXX/events GET プラグインのイベント情報 command getEvents
   /puligins/XXXX/typess GET プラグインのイベント情報 command getTypes
   /puligins/XXXX/config GET プラグインの設定取得 command getConfig
   /puligins/XXXX/config PUT プラグインの設定更新 reload

   -----

   c.engine.GET("/", c.getRoot)
   c.engine.POST("/gox", c.postGox)
   c.engine.GET("/gox/config", c.getGoxConfig)
   c.engine.PUT("/gox/config", c.putGoxConfig)
   c.engine.GET("/plugins", c.getPlugins)
   c.engine.POST("/plugins", c.postPlugins)
   c.engine.DELETE("/plugins/:pluginName", c.deletePlugin)
   c.engine.POST("/plugins/:pluginName", c.postPlugin)
   c.engine.GET("/plugins/:pluginName/help", c.getPluginHelp)
   c.engine.GET("/plugins/:pluginName/commands", c.getPluginCommands)
   c.engine.GET("/plugins/:pluginName/events", c.getPluginEvents)
   c.engine.GET("/plugins/:pluginName/config", c.getPluginConfig)
   c.engine.PUT("/plugins/:pluginName/config", c.putPluginConfig)
*/

func (c *controller) getRoot(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"configPath":    c.configPath,
		"goxHome":       c.goxHome,
		"availablePath": c.availablePath,
		"enablePath":    enablePath,
		"address":       address,
		"port":          port,
	})
}
