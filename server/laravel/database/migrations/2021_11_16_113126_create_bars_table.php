<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateBarsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('bars', function (Blueprint $table) {
            $table->string('id')->primary();;
            $table->string('name');
            $table->string('slug');
            $table->string('descr');
            $table->string('lat');
            $table->string('lon');
            $table->string('city');
            $table->string('address');
            $table->integer('status');
            $table->string('owner');
            $table->timestamps();
        });
    }
    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('bars');
    }
}
