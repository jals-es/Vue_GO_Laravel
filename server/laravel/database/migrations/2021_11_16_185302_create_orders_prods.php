<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateOrdersProds extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('order_prods', function (Blueprint $table) {
            $table->string('id')->primary();
            $table->string('id_order');
            $table->string('name');
            $table->string('descr');
            $table->string('typename');
            $table->string('typedescr');
            $table->float('price');
            $table->integer('status');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('order_prods');
    }
}
