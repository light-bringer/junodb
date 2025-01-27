//
//  Copyright 2023 PayPal Inc.
//
//  Licensed to the Apache Software Foundation (ASF) under one or more
//  contributor license agreements.  See the NOTICE file distributed with
//  this work for additional information regarding copyright ownership.
//  The ASF licenses this file to You under the Apache License, Version 2.0
//  (the "License"); you may not use this file except in compliance with
//  the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//

package cli

import (
	"errors"
	"sync"
)

type IRetryable interface {
	Retryable() bool
}

type Error struct {
	What string
}

var (
	ErrConnect         = &Error{"connection error"}
	ErrResponseTimeout = &Error{"response timeout"}
	rwMutex            sync.RWMutex
)

func specialError(err error) bool {
	if errors.Is(err, ErrConnect) ||
		errors.Is(err, ErrResponseTimeout) {
		return true
	}
	return false
}

func (e *Error) Retryable() bool { return false }

type RetryableError struct {
	What string
}

func (e *RetryableError) Retryable() bool { return true }

func (e *Error) Error() string {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return "error: " + e.What
}

func (e *RetryableError) Error() string {
	return "error: " + e.What
}

func (e *Error) SetError(v string) {
	rwMutex.Lock()
	e.What = v
	rwMutex.Unlock()
}

func NewError(err error) *Error {
	return &Error{
		What: err.Error(),
	}
}

func NewErrorWithString(err string) *Error {
	return &Error{What: err}
}

/*
	// DON'T RETRY,
	//   non-system error, invocation method can handle this response codes
*	case OperationStatus::NoKey:
*	case OperationStatus::DupKey:
	case OperationStatus::DataExpired:
*	case OperationStatus::BadParam:
	case OperationStatus::VersionTooOld:
	case OperationStatus::VersionConflict:
*	case OperationStatus::NoUncommitted:
	case OperationStatus::DuplicateRequest:
	case OperationStatus::NotAppendable:
		return true;

		// RETRY following response codes
*	case OperationStatus::BadMsg:
*	case OperationStatus::OutOfMem:
*	case OperationStatus::NoStorageServer:
	case OperationStatus::StorageServerTimeout:
*	case OperationStatus::RecordLocked:
*	case OperationStatus::BadRequestID:
		// check rethrow flag to figure out if
		// we have to throw exception or return false
		if (_rethrow)
		{
			_trans.AddData(CLIENT_CAL_EXCEPTION, CLIENT_CAL_BADPARAM_EXCEPTION);
			m_tmp_string.copy_formatted("MayflyServer:RetriableSystemError responseStatus='%s'. Response='%s' ",
					OperationStatus::get_err_text(responseStatus),
					_response.to_string().chars());
			_trans.AddData(CLIENT_CAL_DETAILS, m_tmp_string);
			WRITE_LOG_ENTRY(&m_logger, LOG_WARNING, m_tmp_string.chars());

			_trans.SetStatus(CalTransaction::Status(CAL::TRANS_ERROR, CLIENT_CAL_TYPE, CAL::SYS_ERR_INTERNAL, "-1"));
			throw BadParamException(m_tmp_string);
		}
		return false;

		// DONT'T RETRY, system error throw exception
*	case OperationStatus::ServiceDenied:
	case OperationStatus::Inserting:
*/
