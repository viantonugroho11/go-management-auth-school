package server

import (
	"go-management-auth-school/config"

	"github.com/labstack/echo/v4"

	authServices "go-management-auth-school/service/auth"
	authController "go-management-auth-school/controller/auth"

	adminRepository "go-management-auth-school/repository/admin"
	adminServices "go-management-auth-school/service/admin"
	adminController "go-management-auth-school/controller/admin"

	classRepository "go-management-auth-school/repository/class"
	classServices "go-management-auth-school/service/class"
	classController "go-management-auth-school/controller/class"

	

	lessonRepository "go-management-auth-school/repository/lesson"
	lessonServices "go-management-auth-school/service/lesson"
	lessonController "go-management-auth-school/controller/lesson"

	majorRepository "go-management-auth-school/repository/major"
	majorServices "go-management-auth-school/service/major"
	majorController "go-management-auth-school/controller/major"

	mappingCourseRepository "go-management-auth-school/repository/mapping_course"
	mappingCourseServices "go-management-auth-school/service/mapping_course"
	mappingCourseController "go-management-auth-school/controller/mapping_course"

	mappingStudentRepository "go-management-auth-school/repository/mapping_student"
	mappingStudentServices "go-management-auth-school/service/mapping_student"
	mappingStudentController "go-management-auth-school/controller/mapping_student"

	parentRepository "go-management-auth-school/repository/parent"
	parentServices "go-management-auth-school/service/parent"
	parentController "go-management-auth-school/controller/parent"

	userRepository "go-management-auth-school/repository/user"
	userServices "go-management-auth-school/service/user"
	userController "go-management-auth-school/controller/user"

	studentController "go-management-auth-school/controller/student"
	studentRepository "go-management-auth-school/repository/student"
	studentServices "go-management-auth-school/service/student"
)


func InitApp(router *echo.Echo, conf config.Config, unitTest bool) {

	config.MasterDB = config.SetupMasterDB(conf)
	// setup slave db
	config.SlaveDB = config.SetupSlaveDB(conf)



	v1 := router.Group("/v1")


	// v2 := router.Group("/v2")

	// v1 api group
	apiUserV1 := v1.Group("/apiUser")
	apiAdminV1 := v1.Group("/apiAdmin")
	apiAuthV1 := v1.Group("/apiAuth")
	apiStaticv1 := v1.Group("/apiStatic")


	apiUserV1.GET("/health", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})

	apiAdminV1.GET("/test", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})

	apiAuthV1.GET("/test", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})

	apiStaticv1.GET("/test", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})



	adminRepo := adminRepository.NewAdminRepo(config.MasterDB, config.SlaveDB)
	classRepo := classRepository.NewClassRepo(config.MasterDB, config.SlaveDB)
	lessonRepo := lessonRepository.NewLessonRepo(config.MasterDB, config.SlaveDB)
	majorRepo := majorRepository.NewMajorRepo(config.MasterDB, config.SlaveDB)
	mpCourseRepo := mappingCourseRepository.NewMpCourseRepo(config.MasterDB, config.SlaveDB)
	mpStudentRepo := mappingStudentRepository.NewMpStudentRepo(config.MasterDB, config.SlaveDB)
	userRepo := userRepository.NewUserRepo(config.MasterDB, config.SlaveDB)
	parentRepo := parentRepository.NewParentRepo(config.MasterDB, config.SlaveDB)
	studentRepo := studentRepository.NewStudentRepo(config.MasterDB, config.SlaveDB)


	adminService := adminServices.NewAdminService(adminRepo)
	classService := classServices.NewClassService(classRepo)
	lessonService := lessonServices.NewLessonService(lessonRepo)
	majorService := majorServices.NewMajorService(majorRepo)
	mapCourseService := mappingCourseServices.NewMappingCourseService(mpCourseRepo)
	mappingStudentService := mappingStudentServices.NewMappingStudentService(mpStudentRepo)
	userService := userServices.NewUserService(userRepo)
	parentService := parentServices.NewParentService(parentRepo)

	studentService := studentServices.NewStudentService(studentRepo)
	authService := authServices.NewAuthService(userRepo, conf, mapCourseService, studentService, mappingStudentService, userService)


	adminRouter := "/admin"
	adminUserRouter := apiUserV1.Group(adminRouter)
	adminAdminRouter := apiAdminV1.Group(adminRouter)
	adminStaticRouter := apiStaticv1.Group(adminRouter)
	adminAuthRouter := apiAuthV1.Group(adminRouter)
	adminControllers := adminController.NewAdminController(adminService)
	adminControllers.InitializeRoutes(adminUserRouter, adminAdminRouter, adminStaticRouter, adminAuthRouter)

	authRouter := "/auth"
	authUserRouter := apiUserV1.Group(authRouter)
	authAdminRouter := apiAdminV1.Group(authRouter)
	authStaticRouter := apiStaticv1.Group(authRouter)
	authAuthRouter := apiAuthV1.Group(authRouter)
	authControllers := authController.NewAuthController(authService)
	authControllers.InitializeRoutes(authUserRouter, authAdminRouter, authStaticRouter, authAuthRouter)

	classRouter := "/class"
	classUserRouter := apiUserV1.Group(classRouter)
	classAdminRouter := apiAdminV1.Group(classRouter)
	classStaticRouter := apiStaticv1.Group(classRouter)
	classAuthRouter := apiAuthV1.Group(classRouter)
	classControllers := classController.NewClassController(classService)
	classControllers.InitializeRoutes(classUserRouter, classAdminRouter, classStaticRouter, classAuthRouter)

	lessonRouter := "/lesson"
	lessonUserRouter := apiUserV1.Group(lessonRouter)
	lessonAdminRouter := apiAdminV1.Group(lessonRouter)
	lessonStaticRouter := apiStaticv1.Group(lessonRouter)
	lessonAuthRouter := apiAuthV1.Group(lessonRouter)
	lessonControllers := lessonController.NewLessonController(lessonService)
	lessonControllers.InitializeRoutes(lessonUserRouter, lessonAdminRouter, lessonStaticRouter, lessonAuthRouter)

	majorRouter := "/major"
	majorUserRouter := apiUserV1.Group(majorRouter)
	majorAdminRouter := apiAdminV1.Group(majorRouter)
	majorStaticRouter := apiStaticv1.Group(majorRouter)
	majorAuthRouter := apiAuthV1.Group(majorRouter)
	majorControllers := majorController.NewMajorController(majorService)
	majorControllers.InitializeRoutes(majorUserRouter, majorAdminRouter, majorStaticRouter, majorAuthRouter)

	mappingCourseRouter := "/mappingCourse"
	mappingCourseUserRouter := apiUserV1.Group(mappingCourseRouter)
	mappingCourseAdminRouter := apiAdminV1.Group(mappingCourseRouter)
	mappingCourseStaticRouter := apiStaticv1.Group(mappingCourseRouter)
	mappingCourseAuthRouter := apiAuthV1.Group(mappingCourseRouter)
	mappingCourseControllers := mappingCourseController.NewMappingCourseController(mapCourseService)
	mappingCourseControllers.InitializeRoutes(mappingCourseUserRouter, mappingCourseAdminRouter, mappingCourseStaticRouter, mappingCourseAuthRouter)

	mappingStudentRouter := "/mappingStudent"
	mappingStudentUserRouter := apiUserV1.Group(mappingStudentRouter)
	mappingStudentAdminRouter := apiAdminV1.Group(mappingStudentRouter)
	mappingStudentStaticRouter := apiStaticv1.Group(mappingStudentRouter)
	mappingStudentAuthRouter := apiAuthV1.Group(mappingStudentRouter)
	mappingStudentControllers := mappingStudentController.NewMappingStudentController(mappingStudentService)
	mappingStudentControllers.InitializeRoutes(mappingStudentUserRouter, mappingStudentAdminRouter, mappingStudentStaticRouter, mappingStudentAuthRouter)

	parentRouter := "/parent"
	parentUserRouter := apiUserV1.Group(parentRouter)
	parentAdminRouter := apiAdminV1.Group(parentRouter)
	parentStaticRouter := apiStaticv1.Group(parentRouter)
	parentAuthRouter := apiAuthV1.Group(parentRouter)
	parentControllers := parentController.NewParentController(parentService)
	parentControllers.InitializeRoutes(parentUserRouter, parentAdminRouter, parentStaticRouter, parentAuthRouter)

	userRouter := "/user"
	userUserRouter := apiUserV1.Group(userRouter)
	userAdminRouter := apiAdminV1.Group(userRouter)
	userStaticRouter := apiStaticv1.Group(userRouter)
	userAuthRouter := apiAuthV1.Group(userRouter)
	userControllers := userController.NewUserController(userService)
	userControllers.InitializeRoutes(userUserRouter, userAdminRouter, userStaticRouter, userAuthRouter)

	studentRouter := "/student"
	studentUserRouter := apiUserV1.Group(studentRouter)
	studentAdminRouter := apiAdminV1.Group(studentRouter)
	studentStaticRouter := apiStaticv1.Group(studentRouter)
	studentAuthRouter := apiAuthV1.Group(studentRouter)
	studentControllers := studentController.NewStudentController(studentService)
	studentControllers.InitializeRoutes(studentUserRouter, studentAdminRouter, studentStaticRouter, studentAuthRouter)



	// v2 api group
	// apiUserV2 := v2.Group("/apiUser")
	// apiAdminV2 := v2.Group("/apiAdmin")
	// apiAuthV2 := v2.Group("/apiAuth")
	// apiStaticv2 := v2.Group("/apiStatic")


}
