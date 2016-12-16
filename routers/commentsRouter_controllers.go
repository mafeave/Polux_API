package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["Polux_API/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EntidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EntidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EntidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EntidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EntidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:FormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:FormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:FormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:FormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:FormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PaqueteRespuestasController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PaqueteRespuestasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PaqueteRespuestasController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PaqueteRespuestasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PaqueteRespuestasController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PaqueteRespuestasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PaqueteRespuestasController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PaqueteRespuestasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PaqueteRespuestasController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PaqueteRespuestasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RevisionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RevisionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RevisionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RevisionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:RevisionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Polux_API/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["Polux_API/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
