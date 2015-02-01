/*****************************************************************************
 * Copyright (c) 2015 Ratouit Thomas                                    	 *
 *                                                                           *
 * This program is free software; you can redistribute it and/or modify it   *
 * under the terms of the GNU Lesser General Public License as published by  *
 * the Free Software Foundation; either version 3 of the License, or (at     *
 * your option) any later version.                                           *
 *                                                                           *
 * This program is distributed in the hope that it will be useful, but       *
 * WITHOUT ANY WARRANTY; without even the implied warranty of                *
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser   *
 * General Public License for more details.                                  *
 *                                                                           *
 * You should have received a copy of the GNU Lesser General Public License  *
 * along with this program; if not, write to the Free Software Foundation,   *
 * Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA, or go to   *
 * http://www.gnu.org/copyleft/lesser.txt.                                   *
 *****************************************************************************/

/*
The package go-binding-js is a simple Go/Javascript binding for cocos2d-js framework


Here an example from cocos2d-js:

	cc.game.onStart = function(){
	    cc.view.adjustViewPort(true);
	    cc.view.setDesignResolutionSize(800, 450, cc.ResolutionPolicy.SHOW_ALL);
	    cc.view.resizeWithBrowserSize(true);
	    //load resources
	    cc.LoaderScene.preload(g_resources, function () {
	        cc.director.runScene(new HelloWorldScene());
	    }, this);
	};
	cc.game.run();

And the equivalent Go code using the binding:

	package main

	import (
		"github.com/gopherjs/gopherjs/js"
		"github.com/thoratou/go-cocos2d-js/cc"
	)

	func main() {
	cc.Game.OnStart(func() {
		cc.View().AdjustViewPort(true)
		cc.View().SetDesignResolutionSize(800, 450, cc.SHOW_ALL)
		cc.View().ResizeWithBrowserSize(true)
		//load resources
		cc.LoaderScene().Preload(resources, func() {
			cc.Director().RunScene(NewHelloWorldScene())
		})
	})
*/
package gococos2djs
