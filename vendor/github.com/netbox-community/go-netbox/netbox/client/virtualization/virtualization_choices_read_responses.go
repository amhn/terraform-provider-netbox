// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// VirtualizationChoicesReadReader is a Reader for the VirtualizationChoicesRead structure.
type VirtualizationChoicesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationChoicesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationChoicesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationChoicesReadOK creates a VirtualizationChoicesReadOK with default headers values
func NewVirtualizationChoicesReadOK() *VirtualizationChoicesReadOK {
	return &VirtualizationChoicesReadOK{}
}

/*VirtualizationChoicesReadOK handles this case with default header values.

VirtualizationChoicesReadOK virtualization choices read o k
*/
type VirtualizationChoicesReadOK struct {
}

func (o *VirtualizationChoicesReadOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/_choices/{id}/][%d] virtualizationChoicesReadOK ", 200)
}

func (o *VirtualizationChoicesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
