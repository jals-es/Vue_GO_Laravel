<?php

namespace App\Interfaces;

interface ProductRepositoryInterface
{
    public function getAllProducts($barId);
    public function getTypesByProduct($productId);
    public function getExtrasByProduct($productId);
    public function assembleProduct($product, array $types, array $extras);
}
