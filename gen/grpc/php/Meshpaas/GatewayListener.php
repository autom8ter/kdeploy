<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Meshpaas;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * GatewayListener is a single, public tcp listener served by a gateway
 *
 * Generated from protobuf message <code>meshpaas.GatewayListener</code>
 */
class GatewayListener extends \Google\Protobuf\Internal\Message
{
    /**
     * port the gateway listener will listen on ex: 80
     *
     * Generated from protobuf field <code>uint32 port = 1 [(.validator.field) = {</code>
     */
    private $port = 0;
    /**
     * name of the gateway listener ex: http
     *
     * Generated from protobuf field <code>string name = 2 [(.validator.field) = {</code>
     */
    private $name = '';
    /**
     * protocol describes the transport protocol served by this listener
     *
     * Generated from protobuf field <code>.meshpaas.TransportProtocol protocol = 3 [(.validator.field) = {</code>
     */
    private $protocol = 0;
    /**
     * hosts describes hostnames that may route to this listener
     *
     * Generated from protobuf field <code>repeated string hosts = 4 [(.validator.field) = {</code>
     */
    private $hosts;
    /**
     * tls_config provides tls/ssl encryption options
     *
     * Generated from protobuf field <code>.meshpaas.ServerTLSSettings tls_config = 5;</code>
     */
    private $tls_config = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $port
     *           port the gateway listener will listen on ex: 80
     *     @type string $name
     *           name of the gateway listener ex: http
     *     @type int $protocol
     *           protocol describes the transport protocol served by this listener
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $hosts
     *           hosts describes hostnames that may route to this listener
     *     @type \Meshpaas\ServerTLSSettings $tls_config
     *           tls_config provides tls/ssl encryption options
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * port the gateway listener will listen on ex: 80
     *
     * Generated from protobuf field <code>uint32 port = 1 [(.validator.field) = {</code>
     * @return int
     */
    public function getPort()
    {
        return $this->port;
    }

    /**
     * port the gateway listener will listen on ex: 80
     *
     * Generated from protobuf field <code>uint32 port = 1 [(.validator.field) = {</code>
     * @param int $var
     * @return $this
     */
    public function setPort($var)
    {
        GPBUtil::checkUint32($var);
        $this->port = $var;

        return $this;
    }

    /**
     * name of the gateway listener ex: http
     *
     * Generated from protobuf field <code>string name = 2 [(.validator.field) = {</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * name of the gateway listener ex: http
     *
     * Generated from protobuf field <code>string name = 2 [(.validator.field) = {</code>
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
     * protocol describes the transport protocol served by this listener
     *
     * Generated from protobuf field <code>.meshpaas.TransportProtocol protocol = 3 [(.validator.field) = {</code>
     * @return int
     */
    public function getProtocol()
    {
        return $this->protocol;
    }

    /**
     * protocol describes the transport protocol served by this listener
     *
     * Generated from protobuf field <code>.meshpaas.TransportProtocol protocol = 3 [(.validator.field) = {</code>
     * @param int $var
     * @return $this
     */
    public function setProtocol($var)
    {
        GPBUtil::checkEnum($var, \Meshpaas\TransportProtocol::class);
        $this->protocol = $var;

        return $this;
    }

    /**
     * hosts describes hostnames that may route to this listener
     *
     * Generated from protobuf field <code>repeated string hosts = 4 [(.validator.field) = {</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getHosts()
    {
        return $this->hosts;
    }

    /**
     * hosts describes hostnames that may route to this listener
     *
     * Generated from protobuf field <code>repeated string hosts = 4 [(.validator.field) = {</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setHosts($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->hosts = $arr;

        return $this;
    }

    /**
     * tls_config provides tls/ssl encryption options
     *
     * Generated from protobuf field <code>.meshpaas.ServerTLSSettings tls_config = 5;</code>
     * @return \Meshpaas\ServerTLSSettings
     */
    public function getTlsConfig()
    {
        return $this->tls_config;
    }

    /**
     * tls_config provides tls/ssl encryption options
     *
     * Generated from protobuf field <code>.meshpaas.ServerTLSSettings tls_config = 5;</code>
     * @param \Meshpaas\ServerTLSSettings $var
     * @return $this
     */
    public function setTlsConfig($var)
    {
        GPBUtil::checkMessage($var, \Meshpaas\ServerTLSSettings::class);
        $this->tls_config = $var;

        return $this;
    }

}

