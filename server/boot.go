package server

import (
	"go-management-auth-school/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	middlewareCustom "go-management-auth-school/middlewares"

	authController "go-management-auth-school/controller/auth"
	authServices "go-management-auth-school/service/auth"

	adminController "go-management-auth-school/controller/admin"
	adminRepository "go-management-auth-school/repository/admin"
	adminServices "go-management-auth-school/service/admin"

	classController "go-management-auth-school/controller/class"
	classRepository "go-management-auth-school/repository/class"
	classServices "go-management-auth-school/service/class"

	lessonController "go-management-auth-school/controller/lesson"
	lessonRepository "go-management-auth-school/repository/lesson"
	lessonServices "go-management-auth-school/service/lesson"

	majorController "go-management-auth-school/controller/major"
	majorRepository "go-management-auth-school/repository/major"
	majorServices "go-management-auth-school/service/major"

	mappingCourseController "go-management-auth-school/controller/mapping_course"
	mappingCourseRepository "go-management-auth-school/repository/mapping_course"
	mappingCourseServices "go-management-auth-school/service/mapping_course"

	mappingStudentController "go-management-auth-school/controller/mapping_student"
	mappingStudentRepository "go-management-auth-school/repository/mapping_student"
	mappingStudentServices "go-management-auth-school/service/mapping_student"

	parentController "go-management-auth-school/controller/parent"
	parentRepository "go-management-auth-school/repository/parent"
	parentServices "go-management-auth-school/service/parent"

	userController "go-management-auth-school/controller/user"
	userRepository "go-management-auth-school/repository/user"
	userServices "go-management-auth-school/service/user"

	studentController "go-management-auth-school/controller/student"
	studentRepository "go-management-auth-school/repository/student"
	studentServices "go-management-auth-school/service/student"

	teacherController "go-management-auth-school/controller/teacher"
	teacherRepository "go-management-auth-school/repository/teacher"
	teacherServices "go-management-auth-school/service/teacher"

	verifyTokenController "go-management-auth-school/controller/verify_token"
	verifyTokenRepository "go-management-auth-school/repository/verify_token"
	verifyTokenServices "go-management-auth-school/service/verify_token"
)

func InitApp(router *echo.Echo, conf config.Config, unitTest bool) {

	config.MasterDB = config.SetupMasterDB(conf)
	// setup slave db
	config.SlaveDB = config.SetupSlaveDB(conf)

	v1 := router.Group("/v1")

	var _ = middleware.JWTConfig{
		SigningKey: []byte(conf.JwtAuth.JwtSecretKey),
		Claims: &jwt.StandardClaims{},
		// Claims: 	 &authController.JwtCustomClaims{},
	}


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
	teacherRepo := teacherRepository.NewTeacherRepo(config.MasterDB, config.SlaveDB)
	verifyTokenRepo := verifyTokenRepository.NewVerifyTokenRepo(config.MasterDB, config.SlaveDB)

	// service start
	majorService := majorServices.NewMajorService(majorRepo)
	classService := classServices.NewClassService(classRepo)
	teacherService := teacherServices.NewTeacherService(teacherRepo)
	studentService := studentServices.NewStudentService(studentRepo,parentRepo)
	verifyTokenService := verifyTokenServices.NewVerifyTokenService(verifyTokenRepo)

	adminService := adminServices.NewAdminService(adminRepo)
	lessonService := lessonServices.NewLessonService(lessonRepo)
	userService := userServices.NewUserService(userRepo)
	parentService := parentServices.NewParentService(parentRepo, studentService)

	// mapping
	mapCourseService := mappingCourseServices.NewMappingCourseService(mpCourseRepo, lessonService, classService, teacherService)
	mappingStudentService := mappingStudentServices.NewMappingStudentService(mpStudentRepo, studentService, teacherService, classService)
	authService := authServices.NewAuthService(userRepo, conf, mapCourseService, studentService, mappingStudentService, userService, verifyTokenRepo)
	// service end

	// router use
	apiUserV1.Use(middlewareCustom.ValidateToken(conf,verifyTokenRepo))

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

	teacherRouter := "/teacher"
	teacherUserRouter := apiUserV1.Group(teacherRouter)
	teacherAdminRouter := apiAdminV1.Group(teacherRouter)
	teacherStaticRouter := apiStaticv1.Group(teacherRouter)
	teacherAuthRouter := apiAuthV1.Group(teacherRouter)
	teacherControllers := teacherController.NewTeacherController(teacherService)
	teacherControllers.InitializeRoutes(teacherUserRouter, teacherAdminRouter, teacherStaticRouter, teacherAuthRouter)

	verifyTokenRouter := "/verifyToken"
	_ = apiUserV1.Group(verifyTokenRouter)
	_ = apiAdminV1.Group(verifyTokenRouter)
	_ = apiStaticv1.Group(verifyTokenRouter)
	_ = apiAuthV1.Group(verifyTokenRouter)
	_ = verifyTokenController.NewVerifyTokenController(verifyTokenService)
	// verifyTokenControllers.InitializeRoutes(verifyTokenUserRouter, verifyTokenAdminRouter, verifyTokenStaticRouter, verifyTokenAuthRouter)
	// v2 api group
	// apiUserV2 := v2.Group("/apiUser")
	// apiAdminV2 := v2.Group("/apiAdmin")
	// apiAuthV2 := v2.Group("/apiAuth")
	// apiStaticv2 := v2.Group("/apiStatic")

}
