<?php

namespace App\Repositories;

use App\Interfaces\ProductRepositoryInterface;
use App\Models\Product;
use App\Traits\ErrorTrait;

class ProductRepository implements ProductRepositoryInterface
{
    use ErrorTrait;
    public function getAllProducts($barId)
    {
        try {
            $products = Product::where('id_bar', 'LIKE', $barId)->get();
            foreach ($products as $key => $prod) {
                $products[$key]->types = $prod->types;
                $products[$key]->extras = $prod->extras;
            }
            return $products;
        } catch (\Exception $e) {
            return self::throwError($e->getMessage());
        }
    }
}
