<?php

namespace App\Http\Requests;

use Illuminate\Http\Request;

class IncidenceRequest extends Request
{
    public function authorize()
    {
        return true;
    }

    public function rules()
    {
        $rules = [
            'file' => 'required',
            'name' => 'required|max:120',
            'descr' => 'required|max:2048'
        ];

        return $rules;
    }
}
