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

        $request = $this->instance()->all();
        $images = $request['files'];
        $rules = [
            'file' => 'required',
            'name' => 'required|max:120',
            'descr' => 'required|max:2048'
        ];
        foreach($images as $key => $file) {
            $rules['files.'.$key] = 'image|mimes:jpeg,png,gif';
        }
        return $rules;
    }
}
