## C4 Context Diagram

Ниже контекстная диаграмма для монолитного приложения управления отоплением и мониторинга температуры.

```puml
@startuml
title Heating & Temperature Monitoring System - Context

top to bottom direction

!includeurl https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

Person(user, "Пользователь", "Просматривает температуру, управляет отоплением в помещениях")
Person(admin, "Администратор", "Управляет датчиками и настройками системы")

System(system, "Система управления отоплением и мониторинга температуры", "Монолитное приложение на Go с REST API")

System_Ext(sensors, "Датчики температуры", "Передают показания температуры")
System_Ext(client, "Клиентские приложения", "Web/Mobile, используют REST API")

Rel(user, client, "Управляет отоплением, смотрит температуру")
Rel(admin, client, "Администрирует систему и датчики")
Rel(client, system, "REST API")
Rel(sensors, system, "Передают показания")

@enduml
```