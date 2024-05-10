#!/bin/python3.11

import argparse
import logging
import os

logging.basicConfig(level=logging.INFO)
log = logging.getLogger(__name__)

# Replacement takes keys of strings to be replaced and values of environment
# variable names that will be used to replace the key values.
replacement = {
    '<TENANCY_OCID>': os.getenv('OCI_TENANCY_OCID'),
    '<USER_OCID>': os.getenv('OCI_USER_OCID'),
    '<PEM_KEY_FINGERPRINT>': os.getenv('OCI_FINGERPRINT'),
    '<PATH_TO_PRIVATE_KEY>': os.getenv('OCI_PRIVATE_KEY_PATH'),
    '<TENANCY_REGION>': os.getenv('OCI_REGION'),
    '<PRIVATE_KEY_PASSWORD>': os.getenv('OCI_PRIVATE_KEY_PASS', ''),
    '<ENCLOSING_COMPARTMENT_OCID>': os.getenv('OCI_ENCLOSING_CMP', ''),
    '<OCI_COST_CENTER_TAG_OCID>': os.getenv('OCI_CC_TAG', ''),
    '<OCI_ENVIRONMENT_TAG_OCID>': os.getenv('OCI_ENV_TAG', ''),
    '<SECURITY_COMPARTMENT_OCID>': os.getenv('OCI_SEC_CMP', ''),
    '<APPLICATION_COMPARTMENT_OCID>': os.getenv('OCI_APP_CMP', ''),
    '<DATABASE_COMPARTMENT_OCID>': os.getenv('OCI_DB_CMP', ''),
    '<HOME_REGION>': os.getenv('OCI_ID_HOME', ''),
    '<ADMIN_EMAIL>': os.getenv('OCI_ID_ADMIN_EMAIL', ''),
    '<ADMIN_FIRST_NAME>': os.getenv('OCI_ID_ADMIN_FNAME', ''),
    '<ADMIN_LAST_NAME>': os.getenv('OCI_ID_ADMIN_LNAME', ''),
    '<ADMIN_USERNAME>': os.getenv('OCI_ID_ADMIN_USER', ''),
}

def main(args: argparse.Namespace):
    target_file = 'input.auto.tfvars.template'

    validate()

    for root, _, file in os.walk(args.filepath):
        if 'input.auto.tfvars.template' in file:
            path = os.path.abspath(os.path.join(root, target_file))
            log.info(f'Processing {path}')
            process(path)

def process(filepath: os.PathLike):
    try:
        with open(filepath, mode='r+') as f:
            content = f.read()
            for k,v in replacement.items():
                content = content.replace(k, v)
            f.seek(0)
            f.truncate()
            f.write(content)

    except OSError as e:
        log.error(f'Failed to process file {filepath}: {e}')

def validate():
    required_attrs = [
        '<TENANCY_OCID>',
        '<USER_OCID>',
        '<PEM_KEY_FINGERPRINT>',
        '<PATH_TO_PRIVATE_KEY>',
        '<TENANCY_REGION>'
    ]

    for attr in required_attrs:
        if replacement[attr] is None:
            log.critical(f'No environment variable {attr} set - Exiting')
            exit(1)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(prog='preprocessor.py',
        description='Process input files for pipeline consumption')
    parser.add_argument('filepath', help='Root directory to process')
    args = parser.parse_args()
    log.info(f'Parsing file with args {args}')
    main(args)