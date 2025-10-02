//go:build linux && cgo

package main

import "golang.org/x/sys/unix"

/*
#include <sys/errno.h>
#include <stdlib.h>
#include <semaphore.h>

sem_t *go_sem_open(const char *name) {
	return sem_open(name, O_CREAT, 0644, 1);
}
*/
import "C"

func SemOpen(s string) (r uintptr, e error) {
	if r = C.go_sem_open(s); r == 0 {
		e = ErrorCode("open failed", 0)
	}
	return
}

func SemClose(s uintptr) (e error) {
	if r := C.sem_close(s); r != 0 {
		e = ErrorCode("close failed", r)
	}
	return
}

func SemUnlink(s string) (e error) {
	if r := C.sem_unlink(uintptr(CString(s))); r != 0 {
		e = ErrorCode("unlink failed", r)
	}
	return
}

func SemWait(s uintptr) (e error) {
	if r := C.sem_wait(s); r != 0 {
		e = ErrorCode("wait failed", r)
	}
	return
}

func SemTryWait(s uintptr) (e error) {
	if r := C.sem_trywait(s); r != 0 {
		e = ErrorCode("trywait failed", r)
	}
	return
}

func SemPost(s uintptr) (e error) {
	if r := C.sem_post(s); r != 0 {
		e = ErrorCode("post failed", r)
	}
	return
}
