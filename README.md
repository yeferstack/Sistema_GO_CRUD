# Sistema_GO_CRUD
Sistema de Xchango para administrar usuarios, asignar permisos y niveles de acceso. También incluye notificaciones y un historial de acciones de los administradores.

## Especificaciones Técnicas

## Tecnologías Implementadas 
* Golang
* PostgreSQL (PgAdmin 4)

### Variables de Entorno
```shell
INSCRIPCION_CRUD_PGHOST=localhost
INSCRIPCION_CRUD_PGPORT=5432
INSCRIPCION_CRUD_PGUSER=postgres
INSCRIPCION_CRUD_PGPASS=0219251007
INSCRIPCION_CRUD_PGDB=xchango_db
INSCRIPCION_CRUD_PGSCHEMA=Sistema
INSCRIPCION_CRUD_HTTP_PORT=8080
```

### Ejecución del Ejercicio
# 1. Clonar el repositorio
git clone https://github.com/yeferstack/Sistema_GO_CRUD.git

# 2. Entrar a la carpeta del proyecto
cd Sistema_GO_CRUD

# 3. Cambiar a la rama develop
git checkout develop
git pull origin develop

# 4. Configurar variables de entorno y ejecutar
INSCRIPCION_CRUD_PGHOST=localhost \
INSCRIPCION_CRUD_PGPORT=5432 \
INSCRIPCION_CRUD_PGUSER=postgres \
INSCRIPCION_CRUD_PGPASS=0219251007 \
INSCRIPCION_CRUD_PGDB=xchango_db \
INSCRIPCION_CRUD_PGSCHEMA=Sistema \
INSCRIPCION_CRUD_HTTP_PORT=8080 \
bee run