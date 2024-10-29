// Code generated by platformsh/generator/main.go
// DO NOT EDIT

/*
 * Copyright (c) 2021-present Fabien Potencier <fabien@symfony.com>
 *
 * This file is part of Symfony CLI project
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package platformsh

var availablePHPExts = map[string][]string{
	"amqp":            {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"apc":             {"5.4"},
	"apcu":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"apcu_bc":         {"7.0", "7.1", "7.2", "7.3", "7.4"},
	"applepay":        {"7.0", "7.1", "7.3", "7.4"},
	"bcmath":          {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"blackfire":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"bz2":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"calendar":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"ctype":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"curl":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"datadog":         {"8.2", "8.3"},
	"dba":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"dom":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"enchant":         {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"event":           {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"exif":            {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"ffi":             {"7.4", "8.0", "8.1", "8.2", "8.3"},
	"fileinfo":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"ftp":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"gd":              {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"gearman":         {"5.4", "5.5", "5.6"},
	"geoip":           {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"gettext":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"gmp":             {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"gnupg":           {"8.2", "8.3"},
	"http":            {"5.4", "5.5", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"iconv":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"igbinary":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"imagick":         {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0"},
	"imap":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"interbase":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"intl":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"ioncube":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"json":            {"5.6", "7.0", "7.1", "7.2", "7.3", "7.4"},
	"ldap":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"mailparse":       {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"mbstring":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"mcrypt":          {"5.4", "5.5", "5.6", "7.0", "7.1"},
	"memcache":        {"5.4", "5.5", "5.6"},
	"memcached":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"mongo":           {"5.4", "5.5", "5.6"},
	"mongodb":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"msgpack":         {"5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"mssql":           {"5.4", "5.5", "5.6"},
	"mysql":           {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"mysqli":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"mysqlnd":         {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"newrelic":        {"5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"oauth":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"odbc":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"opcache":         {"5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"openswoole":      {"8.2", "8.3"},
	"opentelemetry":   {"8.2", "8.3"},
	"pdo":             {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"pdo_dblib":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"pdo_firebird":    {"5.4", "5.5", "5.6", "7.0", "7.1"},
	"pdo_mysql":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"pdo_odbc":        {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"pdo_pgsql":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"pdo_sqlite":      {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"pdo_sqlsrv":      {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"pecl-http":       {"5.6"},
	"pgsql":           {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"phar":            {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"phpdbg":          {"5.6"},
	"pinba":           {"5.4", "5.5", "5.6"},
	"posix":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"propro":          {"5.6"},
	"protobuf":        {"8.1"},
	"pspell":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"pthreads":        {"7.1", "7.2"},
	"raphf":           {"5.6", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"rdkafka":         {"8.1", "8.2", "8.3"},
	"readline":        {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"recode":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3"},
	"redis":           {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"shmop":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"simplexml":       {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"snmp":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"soap":            {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"sockets":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"sodium":          {"7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"sourceguardian":  {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"spplus":          {"5.4", "5.5"},
	"sqlite3":         {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"sqlsrv":          {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"ssh2":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"swoole":          {"8.2", "8.3"},
	"sybase":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"sysvmsg":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"sysvsem":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"sysvshm":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"tideways":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"tideways_xhprof": {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"tidy":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"tokenizer":       {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"uuid":            {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"wddx":            {"7.0", "7.1", "7.2", "7.3", "7.4"},
	"xcache":          {"5.4", "5.5"},
	"xdebug":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"xhprof":          {"5.4", "5.5", "5.6"},
	"xml":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"xmlreader":       {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"xmlrpc":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"xmlwriter":       {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"xsl":             {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"yaml":            {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
	"zbarcode":        {"7.0", "7.1", "7.2", "7.3"},
	"zendopcache":     {"5.4"},
	"zip":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2", "8.3"},
}

var availableServices = []*service{
	{
		Type: "chrome-headless",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"73", "80", "81", "83", "84", "86", "91", "95", "113", "120"},
		},
	},
	{
		Type: "elasticsearch",
		Versions: serviceVersions{
			Deprecated: []string{"1.4", "1.7", "2.4", "5.2", "5.4", "6.5", "6.8", "7.2", "7.5", "7.7", "7.9", "7.10"},
			Supported:  []string{"7.17", "8.5"},
		},
	},
	{
		Type: "influxdb",
		Versions: serviceVersions{
			Deprecated: []string{"1.2", "1.3", "1.7", "1.8", "2.2"},
			Supported:  []string{"2.3", "2.7"},
		},
	},
	{
		Type: "kafka",
		Versions: serviceVersions{
			Deprecated: []string{"2.1", "2.2", "2.3", "2.4", "2.5", "2.6", "2.7"},
			Supported:  []string{"3.2", "3.4", "3.6", "3.7"},
		},
	},
	{
		Type: "mariadb",
		Versions: serviceVersions{
			Deprecated: []string{"5.5", "10.0", "10.1", "10.2", "10.3"},
			Supported:  []string{"10.4", "10.5", "10.6", "10.11", "11.0", "11.2", "11.4"},
		},
	},
	{
		Type: "memcached",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"1.4", "1.5", "1.6"},
		},
	},
	{
		Type: "mongodb",
		Versions: serviceVersions{
			Deprecated: []string{"3.0", "3.2", "3.4", "3.6", "4.0.3"},
			Supported: []string{},
		},
	},
	{
		Type: "mongodb-enterprise",
		Versions: serviceVersions{
			Deprecated: []string{"4.0", "4.2"},
			Supported:  []string{"4.4", "5.0", "6.0", "7.0"},
		},
	},
	{
		Type: "mysql",
		Versions: serviceVersions{
			Deprecated: []string{"5.5", "10.0", "10.1", "10.2"},
			Supported:  []string{"10.3", "10.4", "10.5", "10.6", "10.11", "11.0"},
		},
	},
	{
		Type: "network-storage",
		Versions: serviceVersions{
			Deprecated: []string{"1.0"},
			Supported:  []string{"2.0"},
		},
	},
	{
		Type: "opensearch",
		Versions: serviceVersions{
			Deprecated: []string{"1.1", "1.2"},
			Supported:  []string{"1", "2"},
		},
	},
	{
		Type: "oracle-mysql",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"5.7", "8.0"},
		},
	},
	{
		Type: "postgresql",
		Versions: serviceVersions{
			Deprecated: []string{"9.3", "9.4", "9.5", "9.6", "10", "11"},
			Supported:  []string{"12", "13", "14", "15", "16"},
		},
	},
	{
		Type: "rabbitmq",
		Versions: serviceVersions{
			Deprecated: []string{"3.5", "3.6", "3.7", "3.8", "3.9", "3.10", "3.11"},
			Supported:  []string{"3.12", "3.13"},
		},
	},
	{
		Type: "redis",
		Versions: serviceVersions{
			Deprecated: []string{"2.8", "3.0", "3.2", "4.0", "5.0", "6.0"},
			Supported:  []string{"6.2", "7.0", "7.2"},
		},
	},
	{
		Type: "solr",
		Versions: serviceVersions{
			Deprecated: []string{"3.6", "4.10", "6.3", "6.6", "7.6", "7.7", "8.0", "8.4", "8.6"},
			Supported:  []string{"8.11", "9.1", "9.2", "9.4", "9.6"},
		},
	},
	{
		Type: "varnish",
		Versions: serviceVersions{
			Deprecated: []string{"5.1", "5.2", "6.3", "6.4", "7.1"},
			Supported:  []string{"6.0", "7.2", "7.3"},
		},
	},
	{
		Type: "vault-kms",
		Versions: serviceVersions{
			Deprecated: []string{"1.6", "1.8"},
			Supported:  []string{"1.12"},
		},
	},
}
