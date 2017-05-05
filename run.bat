genevent.exe

copy F:\genevent\ClientEventType.java D:\idea\zhajinhua_server\game_common\src\main\java\com\netease\game\protobuf\ClientEventType.java /Y
copy F:\genevent\ServerEventType.java D:\idea\zhajinhua_server\game_common\src\main\java\com\netease\game\protobuf\ServerEventType.java /Y
copy F:\genevent\BackendEventType.java D:\idea\zhajinhua_server\game_common\src\main\java\com\netease\game\protobuf\BackendEventType.java /Y

genevent.exe event.xml event_client.tpl

pause
