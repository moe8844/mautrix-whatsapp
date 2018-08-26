// mauLogger - A logger for Go programs
// Copyright (C) 2016-2018 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package maulogger

import (
	"os"
)

// DefaultLogger ...
var DefaultLogger = Create().(*BasicLogger)

// SetWriter formats the given parts with fmt.Sprint and logs the result with the SetWriter level
func SetWriter(w *os.File) {
	DefaultLogger.SetWriter(w)
}

// OpenFile formats the given parts with fmt.Sprint and logs the result with the OpenFile level
func OpenFile() error {
	return DefaultLogger.OpenFile()
}

// Close formats the given parts with fmt.Sprint and logs the result with the Close level
func Close() {
	DefaultLogger.Close()
}

// Sub creates a Sublogger
func Sub(module string) Logger {
	return DefaultLogger.Sub(module)
}

// Raw formats the given parts with fmt.Sprint and logs the result with the Raw level
func Raw(level Level, module, message string) {
	DefaultLogger.Raw(level, module, message)
}

// Log formats the given parts with fmt.Sprint and logs the result with the given level
func Log(level Level, parts ...interface{}) {
	DefaultLogger.DefaultSub.Log(level, parts...)
}

// Logln formats the given parts with fmt.Sprintln and logs the result with the given level
func Logln(level Level, parts ...interface{}) {
	DefaultLogger.DefaultSub.Logln(level, parts...)
}

// Logf formats the given message and args with fmt.Sprintf and logs the result with the given level
func Logf(level Level, message string, args ...interface{}) {
	DefaultLogger.DefaultSub.Logf(level, message, args...)
}

// Logfln formats the given message and args with fmt.Sprintf, appends a newline and logs the result with the given level
func Logfln(level Level, message string, args ...interface{}) {
	DefaultLogger.DefaultSub.Logfln(level, message, args...)
}

// Debug formats the given parts with fmt.Sprint and logs the result with the Debug level
func Debug(parts ...interface{}) {
	DefaultLogger.DefaultSub.Debug(parts...)
}

// Debugln formats the given parts with fmt.Sprintln and logs the result with the Debug level
func Debugln(parts ...interface{}) {
	DefaultLogger.DefaultSub.Debugln(parts...)
}

// Debugf formats the given message and args with fmt.Sprintf and logs the result with the Debug level
func Debugf(message string, args ...interface{}) {
	DefaultLogger.DefaultSub.Debugf(message, args...)
}

// Info formats the given parts with fmt.Sprint and logs the result with the Info level
func Info(parts ...interface{}) {
	DefaultLogger.DefaultSub.Info(parts...)
}

// Infoln formats the given parts with fmt.Sprintln and logs the result with the Info level
func Infoln(parts ...interface{}) {
	DefaultLogger.DefaultSub.Infoln(parts...)
}

// Infof formats the given message and args with fmt.Sprintf and logs the result with the Info level
func Infof(message string, args ...interface{}) {
	DefaultLogger.DefaultSub.Infof(message, args...)
}

// Warn formats the given parts with fmt.Sprint and logs the result with the Warn level
func Warn(parts ...interface{}) {
	DefaultLogger.DefaultSub.Warn(parts...)
}

// Warnln formats the given parts with fmt.Sprintln and logs the result with the Warn level
func Warnln(parts ...interface{}) {
	DefaultLogger.DefaultSub.Warnln(parts...)
}

// Warnf formats the given message and args with fmt.Sprintf and logs the result with the Warn level
func Warnf(message string, args ...interface{}) {
	DefaultLogger.DefaultSub.Warnf(message, args...)
}

// Error formats the given parts with fmt.Sprint and logs the result with the Error level
func Error(parts ...interface{}) {
	DefaultLogger.DefaultSub.Error(parts...)
}

// Errorln formats the given parts with fmt.Sprintln and logs the result with the Error level
func Errorln(parts ...interface{}) {
	DefaultLogger.DefaultSub.Errorln(parts...)
}

// Errorf formats the given message and args with fmt.Sprintf and logs the result with the Error level
func Errorf(message string, args ...interface{}) {
	DefaultLogger.DefaultSub.Errorf(message, args...)
}

// Fatal formats the given parts with fmt.Sprint and logs the result with the Fatal level
func Fatal(parts ...interface{}) {
	DefaultLogger.DefaultSub.Fatal(parts...)
}

// Fatalln formats the given parts with fmt.Sprintln and logs the result with the Fatal level
func Fatalln(parts ...interface{}) {
	DefaultLogger.DefaultSub.Fatalln(parts...)
}

// Fatalf formats the given message and args with fmt.Sprintf and logs the result with the Fatal level
func Fatalf(message string, args ...interface{}) {
	DefaultLogger.DefaultSub.Fatalf(message, args...)
}

// Write formats the given parts with fmt.Sprint and logs the result with the Write level
func (log *BasicLogger) Write(p []byte) (n int, err error) {
	return log.DefaultSub.Write(p)
}

// Log formats the given parts with fmt.Sprint and logs the result with the given level
func (log *BasicLogger) Log(level Level, parts ...interface{}) {
	log.DefaultSub.Log(level, parts...)
}

// Logln formats the given parts with fmt.Sprintln and logs the result with the given level
func (log *BasicLogger) Logln(level Level, parts ...interface{}) {
	log.DefaultSub.Logln(level, parts...)
}

// Logf formats the given message and args with fmt.Sprintf and logs the result with the given level
func (log *BasicLogger) Logf(level Level, message string, args ...interface{}) {
	log.DefaultSub.Logf(level, message, args...)
}

// Logfln formats the given message and args with fmt.Sprintf, appends a newline and logs the result with the given level
func (log *BasicLogger) Logfln(level Level, message string, args ...interface{}) {
	log.DefaultSub.Logfln(level, message, args...)
}

// Debug formats the given parts with fmt.Sprint and logs the result with the Debug level
func (log *BasicLogger) Debug(parts ...interface{}) {
	log.DefaultSub.Debug(parts...)
}

// Debugln formats the given parts with fmt.Sprintln and logs the result with the Debug level
func (log *BasicLogger) Debugln(parts ...interface{}) {
	log.DefaultSub.Debugln(parts...)
}

// Debugf formats the given message and args with fmt.Sprintf and logs the result with the Debug level
func (log *BasicLogger) Debugf(message string, args ...interface{}) {
	log.DefaultSub.Debugf(message, args...)
}

// Debugfln formats the given message and args with fmt.Sprintf, appends a newline and logs the result with the Debug level
func (log *BasicLogger) Debugfln(message string, args ...interface{}) {
	log.DefaultSub.Debugfln(message, args...)
}

// Info formats the given parts with fmt.Sprint and logs the result with the Info level
func (log *BasicLogger) Info(parts ...interface{}) {
	log.DefaultSub.Info(parts...)
}

// Infoln formats the given parts with fmt.Sprintln and logs the result with the Info level
func (log *BasicLogger) Infoln(parts ...interface{}) {
	log.DefaultSub.Infoln(parts...)
}

// Infofln formats the given message and args with fmt.Sprintf, appends a newline and logs the result with the Info level
func (log *BasicLogger) Infofln(message string, args ...interface{}) {
	log.DefaultSub.Infofln(message, args...)
}

// Infof formats the given message and args with fmt.Sprintf and logs the result with the Info level
func (log *BasicLogger) Infof(message string, args ...interface{}) {
	log.DefaultSub.Infof(message, args...)
}

// Warn formats the given parts with fmt.Sprint and logs the result with the Warn level
func (log *BasicLogger) Warn(parts ...interface{}) {
	log.DefaultSub.Warn(parts...)
}

// Warnln formats the given parts with fmt.Sprintln and logs the result with the Warn level
func (log *BasicLogger) Warnln(parts ...interface{}) {
	log.DefaultSub.Warnln(parts...)
}

// Warnfln formats the given message and args with fmt.Sprintf, appends a newline and logs the result with the Warn level
func (log *BasicLogger) Warnfln(message string, args ...interface{}) {
	log.DefaultSub.Warnfln(message, args...)
}

// Warnf formats the given message and args with fmt.Sprintf and logs the result with the Warn level
func (log *BasicLogger) Warnf(message string, args ...interface{}) {
	log.DefaultSub.Warnf(message, args...)
}

// Error formats the given parts with fmt.Sprint and logs the result with the Error level
func (log *BasicLogger) Error(parts ...interface{}) {
	log.DefaultSub.Error(parts...)
}

// Errorln formats the given parts with fmt.Sprintln and logs the result with the Error level
func (log *BasicLogger) Errorln(parts ...interface{}) {
	log.DefaultSub.Errorln(parts...)
}

// Errorf formats the given message and args with fmt.Sprintf and logs the result with the Error level
func (log *BasicLogger) Errorf(message string, args ...interface{}) {
	log.DefaultSub.Errorf(message, args...)
}

// Errorfln formats the given message and args with fmt.Sprintf, appends a newline and logs the result with the Error level
func (log *BasicLogger) Errorfln(message string, args ...interface{}) {
	log.DefaultSub.Errorfln(message, args...)
}

// Fatal formats the given parts with fmt.Sprint and logs the result with the Fatal level
func (log *BasicLogger) Fatal(parts ...interface{}) {
	log.DefaultSub.Fatal(parts...)
}

// Fatalln formats the given parts with fmt.Sprintln and logs the result with the Fatal level
func (log *BasicLogger) Fatalln(parts ...interface{}) {
	log.DefaultSub.Fatalln(parts...)
}

// Fatalf formats the given message and args with fmt.Sprintf and logs the result with the Fatal level
func (log *BasicLogger) Fatalf(message string, args ...interface{}) {
	log.DefaultSub.Fatalf(message, args...)
}

// Fatalfln formats the given message and args with fmt.Sprintf, appends a newline and logs the result with the Fatal level
func (log *BasicLogger) Fatalfln(message string, args ...interface{}) {
	log.DefaultSub.Fatalfln(message, args...)
}
