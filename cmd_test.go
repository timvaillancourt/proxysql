package proxysql

import "testing"

func TestLoadQueryRulesToRuntime(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = LoadQueryRulesToRuntime(db)
	if err != nil {
		t.Error(err)
	}
}

func TestSaveQueryRulesToDisk(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = SaveQueryRulesToDisk(db)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadUserToRuntime(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = LoadUserToRuntime(db)
	if err != nil {
		t.Error(err)
	}
}

func TestSaveUserToDisk(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = SaveUserToDisk(db)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadServerToRunTime(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = LoadServerToRuntime(db)
	if err != nil {
		t.Error(err)
	}
}

func TestSaveServerToDisk(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = SaveServerToDisk(db)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadSchedulerToRuntime(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = LoadSchedulerToRuntime(db)
	if err != nil {
		t.Error(err)
	}
}

func TestSaveSchedulerToDisk(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = SaveSchedulerToDisk(db)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadMySQLVariablesToRunTime(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = LoadMySQLVariablesToRuntime(db)
	if err != nil {
		t.Error(err)
	}
}

func TestSaveMySQLVariablesToDisk(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = SaveMySQLVariablesToDisk(db)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadAdminVariablesToRuntime(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = LoadAdminVariablesToRuntime(db)
	if err != nil {
		t.Error(err)
	}
}

func TestSaveAdminVariablesToDisk(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = SaveAdminVariablesToDisk(db)
	if err != nil {
		t.Error(err)
	}
}