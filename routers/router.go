// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Polux_API/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/trabajo_grado",
			beego.NSInclude(
				&controllers.TrabajoGradoController{},
			),
		),

		beego.NSNamespace("/tipo_documento",
			beego.NSInclude(
				&controllers.TipoDocumentoController{},
			),
		),

		beego.NSNamespace("/modalidad",
			beego.NSInclude(
				&controllers.ModalidadController{},
			),
		),

		beego.NSNamespace("/comentario",
			beego.NSInclude(
				&controllers.ComentarioController{},
			),
		),

		beego.NSNamespace("/correccion",
			beego.NSInclude(
				&controllers.CorreccionController{},
			),
		),

		beego.NSNamespace("/revision",
			beego.NSInclude(
				&controllers.RevisionController{},
			),
		),

		beego.NSNamespace("/area_conocimiento",
			beego.NSInclude(
				&controllers.AreaConocimientoController{},
			),
		),

		beego.NSNamespace("/areas_trabajo_grado",
			beego.NSInclude(
				&controllers.AreasTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/areas_docente",
			beego.NSInclude(
				&controllers.AreasDocenteController{},
			),
		),

		beego.NSNamespace("/vinculacion_docente",
			beego.NSInclude(
				&controllers.VinculacionDocenteController{},
			),
		),

		beego.NSNamespace("/respuesta",
			beego.NSInclude(
				&controllers.RespuestaController{},
			),
		),

		beego.NSNamespace("/socializacion",
			beego.NSInclude(
				&controllers.SocializacionController{},
			),
		),

		beego.NSNamespace("/respuesta_formato",
			beego.NSInclude(
				&controllers.RespuestaFormatoController{},
			),
		),

		beego.NSNamespace("/formato",
			beego.NSInclude(
				&controllers.FormatoController{},
			),
		),

		beego.NSNamespace("/respuesta_evaluacion",
			beego.NSInclude(
				&controllers.RespuestaEvaluacionController{},
			),
		),

		beego.NSNamespace("/pregunta",
			beego.NSInclude(
				&controllers.PreguntaController{},
			),
		),

		beego.NSNamespace("/pregunta_formato",
			beego.NSInclude(
				&controllers.PreguntaFormatoController{},
			),
		),

		beego.NSNamespace("/estudiante_tg",
			beego.NSInclude(
				&controllers.EstudianteTgController{},
			),
		),

		beego.NSNamespace("/estructura_investigacion_tg",
			beego.NSInclude(
				&controllers.EstructuraInvestigacionTgController{},
			),
		),

		beego.NSNamespace("/categoria",
			beego.NSInclude(
				&controllers.CategoriaController{},
			),
		),

		beego.NSNamespace("/asignatura_tg",
			beego.NSInclude(
				&controllers.AsignaturaTgController{},
			),
		),

		beego.NSNamespace("/asignaturas_elegibles",
			beego.NSInclude(
				&controllers.AsignaturasElegiblesController{},
			),
		),

		beego.NSNamespace("/documento",
			beego.NSInclude(
				&controllers.DocumentoController{},
			),
		),

		beego.NSNamespace("/asignatura_inscrita",
			beego.NSInclude(
				&controllers.AsignaturaInscritaController{},
			),
		),

		beego.NSNamespace("/vinculacion_externa",
			beego.NSInclude(
				&controllers.VinculacionExternaController{},
			),
		),

		beego.NSNamespace("/documento_entidad",
			beego.NSInclude(
				&controllers.DocumentoEntidadController{},
			),
		),

		beego.NSNamespace("/evaluacion",
			beego.NSInclude(
				&controllers.EvaluacionController{},
			),
		),

		beego.NSNamespace("/documento_tg",
			beego.NSInclude(
				&controllers.DocumentoTgController{},
			),
		),

		beego.NSNamespace("/entidad",
			beego.NSInclude(
				&controllers.EntidadController{},
			),
		),

		beego.NSNamespace("/tipo_vinculacion",
			beego.NSInclude(
				&controllers.TipoVinculacionController{},
			),
		),

		beego.NSNamespace("/paquete_respuestas",
			beego.NSInclude(
				&controllers.PaqueteRespuestasController{},
			),
		),

		beego.NSNamespace("/formato_evaluacion_carrera",
			beego.NSInclude(
				&controllers.FormatoEvaluacionCarreraController{},
			),
		),

		beego.NSNamespace("/estado_documento",
			beego.NSInclude(
				&controllers.EstadoDocumentoController{},
			),
		),

		beego.NSNamespace("/solicitud_materias",
			beego.NSInclude(
				&controllers.SolicitudMateriasController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
