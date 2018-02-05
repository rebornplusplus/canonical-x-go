// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2018 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package strutil_test

import (
	"bytes"

	. "gopkg.in/check.v1"

	"github.com/snapcore/snapd/strutil"
)

type limitedWriterSuite struct{}

var _ = Suite(&limitedWriterSuite{})

func (s *limitedWriterSuite) TestWriter(c *C) {
	var buffer bytes.Buffer

	w := strutil.NewLimitedWriter(&buffer, 6)

	data := []byte{'a'}
	n, err := w.Write(data)
	c.Assert(err, IsNil)
	c.Assert(n, Equals, 1)
	c.Assert(buffer.Bytes(), DeepEquals, []byte{'a'})

	data = []byte("bcde")
	n, err = w.Write(data)
	c.Assert(err, IsNil)
	c.Assert(n, Equals, 4)

	n, err = w.Write([]byte("xyz"))
	c.Assert(err, IsNil)
	c.Assert(buffer.Bytes(), DeepEquals, []byte("cdexyz"))
	c.Assert(n, Equals, 3)

	n, err = w.Write([]byte("12"))
	c.Assert(err, IsNil)
	c.Assert(buffer.Bytes(), DeepEquals, []byte("exyz12"))
	c.Assert(n, Equals, 2)
}
