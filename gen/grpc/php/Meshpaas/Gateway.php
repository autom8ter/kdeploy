<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Meshpaas;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>meshpaas.Gateway</code>
 */
class Gateway extends \Google\Protobuf\Internal\Message
{
    /**
     * name of the gateway
     *
     * Generated from protobuf field <code>string name = 1;</code>
     */
    private $name = '';
    /**
     * Generated from protobuf field <code>repeated .meshpaas.GatewayListener listeners = 3;</code>
     */
    private $listeners;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *           name of the gateway
     *     @type \Meshpaas\GatewayListener[]|\Google\Protobuf\Internal\RepeatedField $listeners
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * name of the gateway
     *
     * Generated from protobuf field <code>string name = 1;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * name of the gateway
     *
     * Generated from protobuf field <code>string name = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .meshpaas.GatewayListener listeners = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getListeners()
    {
        return $this->listeners;
    }

    /**
     * Generated from protobuf field <code>repeated .meshpaas.GatewayListener listeners = 3;</code>
     * @param \Meshpaas\GatewayListener[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setListeners($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Meshpaas\GatewayListener::class);
        $this->listeners = $arr;

        return $this;
    }

}

