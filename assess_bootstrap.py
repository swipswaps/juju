#!/usr/bin/env python
from __future__ import print_function

from argparse import ArgumentParser
import os.path

from jujupy import (
    EnvJujuClient,
    SimpleEnvironment,
    temp_bootstrap_env,
    until_timeout,
    )
from utility import scoped_environ


def assess_bootstrap(juju, env, debug, region):
    with scoped_environ():
        juju_bin = os.path.dirname(os.path.abspath(juju))
        os.environ['PATH'] = '{}:{}'.format(juju_bin, os.environ['PATH'])
        client = EnvJujuClient.by_version(SimpleEnvironment.from_config(env),
                                          juju, debug)
    if region is not None:
        client.env.config['region'] = region
    with temp_bootstrap_env(client.juju_home, client):
        client.destroy_environment()
        try:
            client.bootstrap()
        except:
            client.destroy_environment()
    try:
        for ignored in until_timeout(30):
            client.get_status(1)
        print('Environment successfully bootstrapped.')
    finally:
        client.destroy_environment()


def parse_args(argv=None):
    parser = ArgumentParser()
    parser.add_argument('juju', help="The Juju client to use.")
    parser.add_argument('env', help="The environment to test with.")
    parser.add_argument('--debug', action="store_true", default=False,
                        help='Use --debug juju logging.')
    parser.add_argument('--region', help='Override environment region.')
    return parser.parse_args(argv)


def main():
    args = parse_args()
    assess_bootstrap(**args.__dict__)


if __name__ == '__main__':
    main()
