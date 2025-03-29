La Liga Tracker API
Descripción
Este proyecto implementa una API RESTful para el seguimiento de partidos de fútbol, permitiendo la gestión completa (CRUD) de partidos y el registro de eventos como goles, tarjetas y tiempo extra.
Tecnologías utilizadas

Backend: Go 1.21 con Gin Framework
Base de datos: PostgreSQL 13
Containerización: Docker y Docker Compose
Frontend: HTML, CSS y JavaScript (página básica de demostración)

Instalación y ejecución
Prerrequisitos

Docker y Docker Compose instalados en tu sistema

Pasos para ejecutar

Clona este repositorio:
bashCopiargit clone https://github.com/Anton17303/Laboratorio-6.git
cd Laboratorio-6

Levanta los contenedores con Docker Compose:
bashCopiardocker-compose up -d

La API estará disponible en http://localhost:8080/api
Puedes abrir el archivo LaLigaTracker.html en tu navegador para probar la interfaz de usuario

Estructura del proyecto
CopiarLaboratorio-6/
├── Dockerfile              # Configuración para construir la imagen Docker
├── docker-compose.yml      # Configuración de servicios (API y PostgreSQL)
├── go.mod                  # Dependencias de Go
├── go.sum                  # Checksums de dependencias
├── main.go                 # Punto de entrada de la aplicación
├── swagger.yaml            # Documentación OpenAPI/Swagger
├── LaLigaTracker.html      # Interfaz de usuario simple
└── Interno/                # Código interno de la aplicación
    ├── database/           # Configuración y conexión a la base de datos
    ├── handlers/           # Manejadores de las rutas HTTP
    └── Modelo/             # Modelos de datos
Endpoints de la API
Operaciones CRUD básicas

GET /api/matches - Listar todos los partidos
GET /api/matches/:id - Obtener un partido específico
POST /api/matches - Crear un nuevo partido
PUT /api/matches/:id - Actualizar un partido existente
DELETE /api/matches/:id - Eliminar un partido

Operaciones especiales (PATCH)

PATCH /api/matches/:id/goals - Registrar un gol
PATCH /api/matches/:id/yellowcards - Registrar una tarjeta amarilla
PATCH /api/matches/:id/redcards - Registrar una tarjeta roja
PATCH /api/matches/:id/extratime - Establecer tiempo extra

Modelo de datos
goCopiartype Match struct {
	gorm.Model
	HomeTeam    string    `json:"homeTeam"`
	AwayTeam    string    `json:"awayTeam"`
	MatchDate   time.Time `json:"matchDate"`
	Goals       int       `json:"goals"`
	YellowCards int       `json:"yellowCards"`
	RedCards    int       `json:"redCards"`
	ExtraTime   bool      `json:"extraTime"`
}
Colección de Postman/Hoppscotch
Para probar la API, puedes usar la siguiente colección en Hoppscotch:
https://hoppscotch.io/collection/import?data=eyJuYW1lIjoiTGEgTGlnYSBUcmFja2VyIEFQSSIsInZlcnNpb24iOjEsInByb3BzIjp7ImNvbGxlY3Rpb25fbmFtZSI6IkxhIExpZ2EgVHJhY2tlciBBUEkifSwiZGF0YSI6W3sibmFtZSI6Ikxpc3RhciBQYXJ0aWRvcyIsInByb3BzIjp7Im1ldGhvZCI6IkdFVCIsInBhdGgiOiJodHRwOi8vbG9jYWxob3N0OjgwODAvYXBpL21hdGNoZXMiLCJoZWFkZXJzIjpbXSwicGFyYW1zIjpbXSwiYm9keSI6eyJ0eXBlIjoibm9uZSJ9fX0seyJuYW1lIjoiQnVzY2FyIFBhcnRpZG8gcG9yIElEIiwicHJvcHMiOnsibWV0aG9kIjoiR0VUIiwicGF0aCI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MC9hcGkvbWF0Y2hlcy8xIiwiaGVhZGVycyI6W10sInBhcmFtcyI6W10sImJvZHkiOnsidHlwZSI6Im5vbmUifX19LHsibmFtZSI6IkNyZWFyIFBhcnRpZG8iLCJwcm9wcyI6eyJtZXRob2QiOiJQT1NUIiwicGF0aCI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MC9hcGkvbWF0Y2hlcyIsImhlYWRlcnMiOlt7ImtleSI6IkNvbnRlbnQtVHlwZSIsInZhbHVlIjoiYXBwbGljYXRpb24vanNvbiJ9XSwicGFyYW1zIjpbXSwiYm9keSI6eyJ0eXBlIjoianNvbiIsInJhdyI6IntcbiAgXCJob21lVGVhbVwiOiBcIkJhcmNlbG9uYVwiLFxuICBcImF3YXlUZWFtXCI6IFwiUmVhbCBNYWRyaWRcIixcbiAgXCJtYXRjaERhdGVcIjogXCIyMDI1LTAzLTI4VDIwOjAwOjAwWlwiXG59In19fSx7Im5hbWUiOiJBY3R1YWxpemFyIFBhcnRpZG8iLCJwcm9wcyI6eyJtZXRob2QiOiJQVVQiLCJwYXRoIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL2FwaS9tYXRjaGVzLzEiLCJoZWFkZXJzIjpbeyJrZXkiOiJDb250ZW50LVR5cGUiLCJ2YWx1ZSI6ImFwcGxpY2F0aW9uL2pzb24ifV0sInBhcmFtcyI6W10sImJvZHkiOnsidHlwZSI6Impzb24iLCJyYXciOiJ7XG4gIFwiaG9tZVRlYW1cIjogXCJCYXJjZWxvbmFcIixcbiAgXCJhd2F5VGVhbVwiOiBcIlZhbGVuY2lhXCIsXG4gIFwibWF0Y2hEYXRlXCI6IFwiMjAyNS0wNC0wMlQxOTowMDowMFpcIlxufSJ9fX0seyJuYW1lIjoiRWxpbWluYXIgUGFydGlkbyIsInByb3BzIjp7Im1ldGhvZCI6IkRFTEVURSIsInBhdGgiOiJodHRwOi8vbG9jYWxob3N0OjgwODAvYXBpL21hdGNoZXMvMSIsImhlYWRlcnMiOltdLCJwYXJhbXMiOltdLCJib2R5Ijp7InR5cGUiOiJub25lIn19fSx7Im5hbWUiOiJSZWdpc3RyYXIgR29sIiwicHJvcHMiOnsibWV0aG9kIjoiUEFUQ0giLCJwYXRoIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL2FwaS9tYXRjaGVzLzEvZ29hbHMiLCJoZWFkZXJzIjpbXSwicGFyYW1zIjpbXSwiYm9keSI6eyJ0eXBlIjoianNvbiIsInJhdyI6Int9In19fSx7Im5hbWUiOiJSZWdpc3RyYXIgVGFyamV0YSBBbWFyaWxsYSIsInByb3BzIjp7Im1ldGhvZCI6IlBBVENIIiwicGF0aCI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MC9hcGkvbWF0Y2hlcy8xL3llbGxvd2NhcmRzIiwiaGVhZGVycyI6W10sInBhcmFtcyI6W10sImJvZHkiOnsidHlwZSI6Impzb24iLCJyYXciOiJ7fSJ9fX0seyJuYW1lIjoiUmVnaXN0cmFyIFRhcmpldGEgUm9qYSIsInByb3BzIjp7Im1ldGhvZCI6IlBBVENIIiwicGF0aCI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MC9hcGkvbWF0Y2hlcy8xL3JlZGNhcmRzIiwiaGVhZGVycyI6W10sInBhcmFtcyI6W10sImJvZHkiOnsidHlwZSI6Impzb24iLCJyYXciOiJ7fSJ9fX0seyJuYW1lIjoiRXN0YWJsZWNlciBUaWVtcG8gRXh0cmEiLCJwcm9wcyI6eyJtZXRob2QiOiJQQVRDSCIsInBhdGgiOiJodHRwOi8vbG9jYWxob3N0OjgwODAvYXBpL21hdGNoZXMvMS9leHRyYXRpbWUiLCJoZWFkZXJzIjpbXSwicGFyYW1zIjpbXSwiYm9keSI6eyJ0eXBlIjoianNvbiIsInJhdyI6Int9In19fV19