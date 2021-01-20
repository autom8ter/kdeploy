<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Meshpaas;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * AuthzSource is the source of a request
 *
 * Generated from protobuf message <code>meshpaas.AuthzSource</code>
 */
class AuthzSource extends \Google\Protobuf\Internal\Message
{
    /**
     * allow_namespaces restricts access to traffic coming from a particular namespace
     *
     * Generated from protobuf field <code>repeated string allow_namespaces = 1;</code>
     */
    private $allow_namespaces;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $allow_namespaces
     *           allow_namespaces restricts access to traffic coming from a particular namespace
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * allow_namespaces restricts access to traffic coming from a particular namespace
     *
     * Generated from protobuf field <code>repeated string allow_namespaces = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAllowNamespaces()
    {
        return $this->allow_namespaces;
    }

    /**
     * allow_namespaces restricts access to traffic coming from a particular namespace
     *
     * Generated from protobuf field <code>repeated string allow_namespaces = 1;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAllowNamespaces($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->allow_namespaces = $arr;

        return $this;
    }

}

