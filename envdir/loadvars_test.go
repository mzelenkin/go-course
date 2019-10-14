package main

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestLoad(t *testing.T) {
	var err error

	// Создадим рабочую директорию
	dir := os.TempDir()
	envDir := path.Join(dir, "testenv")
	err = os.Mkdir(envDir, 0755)
	if err != nil {
		t.Error("mkdir error: ", err)
	}

	// Создадим внутри первый файл
	filename1 := path.Join(envDir, "VAR1")
	err = ioutil.WriteFile(filename1, []byte("TESTVALUE 1"), 0644)
	if err != nil {
		t.Error(err)
	}

	// и второй
	filename2 := path.Join(envDir, "VAR2")
	err = ioutil.WriteFile(filename2, []byte("TESTVALUE_2"), 0644)
	if err != nil {
		t.Error(err)
	}

	// Пробуем загрузить значение в переменные
	result, err := loadEnv(envDir)
	if err != nil {
		t.Error("LoadEnv error: ", err)
	}

	// Проверка результатов
	if val, ok := result["VAR1"]; !ok || val != "TESTVALUE 1" {
		t.Fail()
	}

	if val, ok := result["VAR2"]; !ok || val != "TESTVALUE_2" {
		t.Fail()
	}

	// Убираем за собой
	err = os.Remove(filename1)
	if err != nil {
		t.Error(err)
	}

	err = os.Remove(filename2)
	if err != nil {
		t.Error(err)
	}

	err = os.Remove(envDir)
	if err != nil {
		t.Error(err)
	}
}
