# -*- coding: utf-8 -*-
# ------------------------------------------------------------------------------
#
#   Copyright 2018-2019 Fetch.AI Limited
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.
#
# ------------------------------------------------------------------------------
"""This test module contains the tests for the `aea add contract` sub-command."""

from unittest import TestCase, mock

from aea.cli import cli

from tests.common.click_testing import CliRunner
from tests.conftest import CLI_LOG_OPTION


@mock.patch("aea.cli.common.try_to_load_agent_config")
@mock.patch("aea.cli.common._validate_config_consistency")
class AddContractCommandTestCase(TestCase):
    """Test that the command 'aea add contract' works as expected."""

    def setUp(self):
        """Set the test up."""
        self.runner = CliRunner()

    @mock.patch("aea.cli.add._add_item")
    def test_add_contract_positive(self, *mocks):
        """Test add contract command positive result."""
        # result = self.runner.invoke(
        #     cli,
        #     [*CLI_LOG_OPTION, "add", "contract", "author/name:0.1.0"],
        #     standalone_mode=False,
        # )
        # self.assertEqual(result.exit_code, 0)
        result = self.runner.invoke(
            cli,
            [*CLI_LOG_OPTION, "add", "--local", "contract", "author/name:0.1.0"],
            standalone_mode=False,
        )
        self.assertEqual(result.exit_code, 0)
