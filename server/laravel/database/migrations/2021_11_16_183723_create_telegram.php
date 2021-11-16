<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateTelegram extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void->primary();
     */
    public function up()
    {
        Schema::create('telegram', function (Blueprint $table) {
            $table->string('id_user');
            $table->string('id_telegram');
            $table->integer('status');
            $table->primary(array('id_user', 'id_telegram'));



        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('telegram');
    }
}
