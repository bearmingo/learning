#!/usr/bin/env python

"""
Tools functions for download or load mnist data
Reference:
https://github.com/datapythonista/mnist/blob/master/mnist/
"""

import os
import shutil
import gzip
import struct
import array
import functools
import operator

import numpy as np
import requests
import fire


SCRIPT_DIR = os.path.abspath(os.path.dirname(__file__))
DATA_DIR = os.path.join(SCRIPT_DIR, 'data')

SITE_URL = 'http://yann.lecun.com/exdb/mnist/'

TRAINING_SET_IMAGES_NAME = 'train-images-idx3-ubyte.gz'
TRAINING_SET_LABELS_NAME = 'train-labels-idx1-ubyte.gz'
TEST_SET_IMAGES_NAME = 't10k-images-idx3-ubyte.gz'
TEST_SET_LABLES_NAME = 't10k-labels-idx1-ubyte.gz'

def download_set(url, localdir):
    "Download from"
    r = requests.get(url, stream=True)
    with open(localdir, 'wb') as tempfile:
        for chunk in r.iter_content(1024):
            tempfile.write(chunk)

def download(target_dir=None):
    if target_dir is None:
        target_dir = DATA_DIR
    tasks = (
        TRAINING_SET_IMAGES_NAME,
        TRAINING_SET_LABELS_NAME,
        TEST_SET_IMAGES_NAME,
        TEST_SET_LABLES_NAME)

    if os.path.exists(target_dir):
        shutil.rmtree(target_dir)
    os.makedirs(target_dir)

    for name in tasks:
        url = SITE_URL + name
        print("Downloading from ", url)
        download_set(url, os.path.join(target_dir, name))


class IdxDecodeError(ValueError):
    """Raised when an invalid idx file is parsed."""
    pass


def parse_idx(fd):
    """Parse an IDX file, and return it as a numpy array.

    Parameters:
    -------
    fd : file
        File descriptor fo the IDX file to parse
    """
    DATA_TYPES = { 0x08: 'B',  # unsigned byte
                   0x09: 'b',  # signed byte
                   0x0b: 'h',  # short (2 bytes)
                   0x0c: 'i',  # int (4 bytes)
                   0x0d: 'f',  # float (4 bytes)
                   0x0e: 'd'}  # double (8 bytes)
    header = fd.read(4)
    if len(header) != 4:
        raise IdxDecodeError(
            'Invalid IDX file, file empty or does not contain a full header')
    
    # > : big-endian
    # H : unsigned short
    # B : unsigned char
    zeros, data_type, num_dimensions = struct.unpack('>HBB', header)
    if zeros != 0:
        raise IdxDecodeError(
            'Invalid IDX file, file must start with two zero bytes.'
            'Found 0x%02x' % zeros)
    
    try:
        data_type = DATA_TYPES[data_type]
    except KeyError:
        raise IdxDecodeError(
            'Unknow data type, '
            '0x%02x in IDX file' % data_type)

    # I : unsigned int
    dimension_size = struct.unpack('>' + 'I' * num_dimensions,
                                   fd.read(4 * num_dimensions))
    data = array.array(data_type, fd.read())
    data.byteswap()  # looks like array.array reads data as little endian

    expected_items = functools.reduce(operator.mul, dimension_size)
    if len(data) != expected_items:
        raise IdxDecodeError('IDX file has wrong number of items. '
                             'Expected: %d. Found: %d' % (expected_items,
                                                          len(data)))

    return np.array(data).reshape(dimension_size)


def parse_mnist_file(fname, target_dir=None, force=False):
    """Parse mnist"""
    fopen = gzip.open if os.path.splitext(fname)[1] == '.gz' else open
    with fopen(fname, 'rb') as fd:
        return parse_idx(fd)


def train_images():
    return parse_mnist_file(os.path.join(DATA_DIR, TRAINING_SET_IMAGES_NAME))


def train_labels():
    return parse_mnist_file(os.path.join(DATA_DIR, TRAINING_SET_LABELS_NAME))


def test_images():
    return parse_mnist_file(os.path.join(DATA_DIR, TEST_SET_IMAGES_NAME))


def test_labels():
    return parse_mnist_file(os.path.join(DATA_DIR, TEST_SET_LABLES_NAME))


def load_data():
    return (train_images(), train_labels(), test_images(), test_labels())


class Helper(object):

    def download(self):
        download()

    def cleanup(self):
        if os.path.exists(DATA_DIR):
            shutil.rmtree(DATA_DIR)

    def test(self):
        train_images()
        train_labels()
        test_images()
        test_labels()

if __name__ == "__main__":
    fire.Fire(Helper)