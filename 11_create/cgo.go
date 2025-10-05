//go:build linux && cgo

package main

import (
	"errors"
	"fmt"
	"unsafe"
)

/*
#include <fcntl.h>
#include <sys/errno.h>
#include <stdlib.h>
#include <semaphore.h>

sem_t *go_sem_open(const char *name) {
	return sem_open(name, O_CREAT, 0644, 1);
}
*/
import "C"

func sem_t_to_uintptr(s *C.sem_t) uintptr {
	return uintptr(unsafe.Pointer(s))
}

func unitptr_to_sem_t(p uintptr) *C.sem_t {
	return (*C.sem_t)(unsafe.Pointer(p))
}

func SemOpen(s string) (r uintptr, e error) {
	if r = sem_t_to_uintptr(C.go_sem_open(C.CString(s))); r == 0 {
		e = errors.New("open failed")
	}
	return
}

func SemClose(s uintptr) (e error) {
	if r := C.sem_close(unitptr_to_sem_t(s)); r != 0 {
		e = errors.New(fmt.Sprint("close failed: %v", r))
	}
	return
}

func SemUnlink(s string) (e error) {
	if r := C.sem_unlink(C.CString(s)); r != 0 {
		e = errors.New(fmt.Sprint("unlink failed: %v", r))
	}
	return
}

func SemWait(s uintptr) (e error) {
	if r := C.sem_wait(unitptr_to_sem_t(s)); r != 0 {
		e = errors.New(fmt.Sprint("wait failed: %v", r))
	}
	return
}

func SemTryWait(s uintptr) (e error) {
	if r := C.sem_trywait(unitptr_to_sem_t(s)); r != 0 {
		e = errors.New(fmt.Sprint("trywait failed: %v", r))
	}
	return
}

func SemPost(s uintptr) (e error) {
	if r := C.sem_post(unitptr_to_sem_t(s)); r != 0 {
		e = errors.New(fmt.Sprint("post failed: %v", r))
	}
	return
}
