# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.

"""
Auto-generated class for PlainObject
"""
import capnp
import os

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class PlainObject(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(**kwargs):
        """
        :type obj: dict
        :rtype: PlainObject
        """

        return PlainObject(**kwargs)

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'PlainObject'
        data = json or kwargs

        # set attributes
        data_types = [dict]
        self.obj = client_support.set_property('obj', data, data_types, False, [], False, True, class_name)

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)

    def to_capnp(self):
        """
        Load the class in capnp schema PlainObject.capnp
        :rtype bytes
        """
        template = capnp.load('%s/PlainObject.capnp' % dir)
        return template.PlainObject.new_message(**self.as_dict()).to_bytes()


class PlainObjectCollection:
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(binary=None):
        """
        Load the binary of PlainObject.capnp into class PlainObject
        :type binary: bytes. If none creates an empty capnp object.
        rtype: PlainObject
        """
        template = capnp.load('%s/PlainObject.capnp' % dir)
        struct = template.PlainObject.from_bytes(binary) if binary else template.PlainObject.new_message()
        return PlainObject(**struct.to_dict(verbose=True))
