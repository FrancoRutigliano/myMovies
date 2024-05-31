<h3 align="center">
  <img src="assets/logo.jpg" width="300" alt="Logo"/><br/>
</h3>

<div align="center"><i>Vive la magia del cine con una reserva de entradas sin complicaciones.</i></div>

## ABOUT
Esta aplicación...

## Get Started
1.Clona el repositorio.
  1.1 ```
2.Clona `.example.env` en `.developemnt.env`
3.Ejecuta el comando `make run-dev` para ejecutar la aplicación en modo desarrollo.
4.Y listo :)

## Features 

### Autenticación de Usuarios:
- [x] Implementar endpoint para registro de nuevos usuarios.
- [x] Implementar endpoint para inicio de sesión de usuarios existentes.
- [X] Implementar sistema de gestión de tokens JWT para autenticación segura.

### Gestión de Roles y Permisos:
- [x] Definir roles necesarios para el sistema (por ejemplo, administrador, usuario regular, etc.).
- [x] Asignar permisos a cada rol para restringir el acceso a las funcionalidades de la API.
- [ ] Implementar lógica de verificación de permisos en los endpoints relevantes.

### Gestión de Salas de Cine:
- [ ] Crear endpoints para la creación, listado, actualización y eliminación de salas de cine.
- [ ] Implementar paginación en el listado de salas de cine.
- [ ] Validar datos de entrada para evitar inconsistencias en la información de las salas de cine.

### Gestión de Películas:
- [ ] Desarrollar endpoints para la creación, listado, actualización y eliminación de películas.
- [ ] Agregar paginación al listado de películas.
- [ ] Permitir la carga de imágenes de portada de películas.

### Gestión de Tickets:
- [ ] Implementar endpoint para la compra de tickets por parte de los usuarios.
- [ ] Desarrollar función para listar tickets adquiridos por un usuario específico.
- [ ] Implementar lógica para cancelar tickets según las políticas establecidas.

