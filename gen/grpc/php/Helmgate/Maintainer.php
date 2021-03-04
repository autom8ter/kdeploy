<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Helmgate;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Maintainer is the maintainer of a helm chart
 *
 * Generated from protobuf message <code>helmgate.Maintainer</code>
 */
class Maintainer extends \Google\Protobuf\Internal\Message
{
    /**
     * name is the name of the maintainer
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     */
    private $name = '';
    /**
     * email is the email of the maintainer
     *
     * Generated from protobuf field <code>string email = 2 [(.validator.field) = {</code>
     */
    private $email = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *           name is the name of the maintainer
     *     @type string $email
     *           email is the email of the maintainer
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * name is the name of the maintainer
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * name is the name of the maintainer
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
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
     * email is the email of the maintainer
     *
     * Generated from protobuf field <code>string email = 2 [(.validator.field) = {</code>
     * @return string
     */
    public function getEmail()
    {
        return $this->email;
    }

    /**
     * email is the email of the maintainer
     *
     * Generated from protobuf field <code>string email = 2 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setEmail($var)
    {
        GPBUtil::checkString($var, True);
        $this->email = $var;

        return $this;
    }

}
