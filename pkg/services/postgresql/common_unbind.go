package postgresql

import (
	"fmt"
)

func unbind(
	enforceSSL bool,
	administratorLogin string,
	serverName string,
	administratorLoginPassword string,
	fullyQualifiedDomainName string,
	loginName string,
) error {
	db, err := getDBConnection(
		enforceSSL,
		administratorLogin,
		serverName,
		administratorLoginPassword,
		fullyQualifiedDomainName,
		primaryDB,
	)
	if err != nil {
		return err
	}
	defer db.Close() // nolint: errcheck

	_, err = db.Exec(
		fmt.Sprintf("drop role %s", loginName),
	)
	if err != nil {
		return fmt.Errorf(`error dropping role "%s": %s`, loginName, err)
	}
	return nil
}
