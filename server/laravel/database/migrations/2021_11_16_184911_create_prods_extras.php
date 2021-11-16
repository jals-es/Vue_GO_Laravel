<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateProdsExtras extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('prods_extras', function (Blueprint $table) {
            $table->string('id')->primary();
            $table->string('id_prod');
            $table->string('name');
            $table->string('descr');
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
        Schema::dropIfExists('prods_extras');
    }
}
