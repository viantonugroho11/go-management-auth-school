package verify_token

import (
	verifyTokenEntity "go-management-auth-school/entity/verify_token"
)

var (

	InsertVerifyToken = `INSERT INTO `+verifyTokenEntity.TableName+` (id,identity_id, token, expired_at) VALUES (?,?, ?, ?)`

	DeleteVerifyToken = `UPDATE `+verifyTokenEntity.TableName+` SET deleted_at = NOW() WHERE identity_id = ?`
)