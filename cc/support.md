Here is the current development status for the binding:

Supported cocos2d version: 3.2 or higher

Based-file support:

|File                                                 | Support         | Comment                           |
|-----------------------------------------------------|-----------------|-----------------------------------|
|CCBoot.js                                            | partial         | TODO: Path, Loader, Windows Events, Sys init |
|cocos2d/actions/CCAction.js                          | partial         | TODO: FiniteTimeAction, Speed, Follow |
|cocos2d/actions/CCActionCatmullRom.js                | not yet         |                                   |
|cocos2d/actions/CCActionEase.js                      | not yet         |                                   |
|cocos2d/actions/CCActionInstant.js                   | not yet         |                                   |
|cocos2d/actions/CCActionInterval.js                  | not yet         |                                   |
|cocos2d/actions/CCActionTween.js                     | not yet         |                                   |
|cocos2d/actions3d/CCActionGrid.js                    | not yet         |                                   |
|cocos2d/actions3d/CCActionGrid3D.js                  | not yet         |                                   |
|cocos2d/actions3d/CCActionPageTurn3D.js              | not yet         |                                   |
|cocos2d/actions3d/CCActionTiledGrid.js               | not yet         |                                   |
|cocos2d/audio/CCAudio.js                             | not yet         |                                   |
|cocos2d/clipping-nodes/CCClippingNode.js             | not yet         |                                   |
|cocos2d/core/base-nodes/CCAtlasNode.js               | not yet         |                                   |
|cocos2d/core/base-nodes/CCNode.js                    | not yet         |                                   |
|cocos2d/core/CCActionManager.js                      | full            |                                   |
|cocos2d/core/CCCamera.js                             | not yet         |                                   |
|cocos2d/core/CCConfiguration.js                      | not yet         |                                   |
|cocos2d/core/CCDirector.js                           | partial         | Dependences: DirectorDelegate, Viewport |
|cocos2d/core/CCDrawingPrimitivesCanvas.js            | not yet         |                                   |
|cocos2d/core/CCScheduler.js                          | partial         | TODO: ListEntry, HashUpdateEntry, HashTimerEntry, Timer |
|cocos2d/core/cocoa/CCAffineTransform.js              | not yet         |                                   |
|cocos2d/core/cocoa/CCGeometry.js                     | full            |                                   |
|cocos2d/core/event-manager/CCEvent.js                | not yet         |                                   |
|cocos2d/core/event-manager/CCEventExtension.js       | not yet         |                                   |
|cocos2d/core/event-manager/CCEventListener.js        | not yet         |                                   |
|cocos2d/core/event-manager/CCEventManager.js         | not yet         |                                   |
|cocos2d/core/event-manager/CCTouch.js                | full            |                                   |
|cocos2d/core/labelttf/CCLabelTTF.js                  | not yet         |                                   |
|cocos2d/core/layers/CCLayer.js                       | partial         | TODO: LayerColor, LayerGradient, LayerMultiplex |
|cocos2d/core/platform/CCClass.js                     | no              | emulates this._super(), no CCClass no go side |
|cocos2d/core/platform/CCCommon.js                    | not yet         |                                   |
|cocos2d/core/platform/CCConfig.js                    | not yet         |                                   |
|cocos2d/core/platform/CCEGLView.js                   | full            |                                   |
|cocos2d/core/platform/CCInputExtension.js            | not yet         |                                   |
|cocos2d/core/platform/CCInputManager.js              | not yet         |                                   |
|cocos2d/core/platform/CCMacro.js                     | not yet         |                                   |
|cocos2d/core/platform/CCSAXParser.js                 | not yet         |                                   |
|cocos2d/core/platform/CCScreen.js                    | not yet         |                                   |
|cocos2d/core/platform/CCTypes.js                     | partial         | TODO: Acceleration, Vertex2F, Vertex3F, Tex2F, BlendFunc, FontDefinition |
|cocos2d/core/platform/CCVisibleRect.js               | not yet         |                                   |
|cocos2d/core/platform/miniFramework.js               | not yet         |                                   |
|cocos2d/core/scenes/CCLoaderScene.js                 | full            |                                   |
|cocos2d/core/scenes/CCScene.js                       | not yet         |                                   |
|cocos2d/core/sprites/CCAnimation.js                  | not yet         |                                   |
|cocos2d/core/sprites/CCAnimationCache.js             | not yet         |                                   |
|cocos2d/core/sprites/CCSprite.js                     | not yet         |                                   |
|cocos2d/core/sprites/CCSpriteBatchNode.js            | not yet         |                                   |
|cocos2d/core/sprites/CCSpriteFrame.js                | not yet         |                                   |
|cocos2d/core/sprites/CCSpriteFrameCache.js           | not yet         |                                   |
|cocos2d/core/support/CCPointExtension.js             | not yet         |                                   |
|cocos2d/core/support/CCVertex.js                     | not yet         |                                   |
|cocos2d/core/support/TransformUtils.js               | not yet         |                                   |
|cocos2d/core/textures/CCTexture2D.js                 | not yet         |                                   |
|cocos2d/core/textures/CCTextureAtlas.js              | not yet         |                                   |
|cocos2d/core/textures/CCTextureCache.js              | not yet         |                                   |
|cocos2d/effects/CCGrabber.js                         | not yet         |                                   |
|cocos2d/effects/CCGrid.js                            | not yet         |                                   |
|cocos2d/labels/CCLabelAtlas.js                       | not yet         |                                   |
|cocos2d/labels/CCLabelBMFont.js                      | not yet         |                                   |
|cocos2d/menus/CCMenu.js                              | full            |                                   |
|cocos2d/menus/CCMenuItem.js                          | partial         | TODO: MenuItemLabel, MenuItemAtlasFont, MenuItemFont, MenuItemSprite, MenuItemImage, |
|cocos2d/motion-streak/CCMotionStreak.js              | not yet         |                                   |
|cocos2d/node-grid/CCNodeGrid.js                      | full            |                                   |
|cocos2d/parallax/CCParallaxNode.js                   | not yet         |                                   |
|cocos2d/particle/CCParticleBatchNode.js              | not yet         |                                   |
|cocos2d/particle/CCParticleExamples.js               | not yet         |                                   |
|cocos2d/particle/CCParticleSystem.js                 | not yet         |                                   |
|cocos2d/progress-timer/CCActionProgressTimer.js      | not yet         |                                   |
|cocos2d/progress-timer/CCProgressTimer.js            | not yet         |                                   |
|cocos2d/render-texture/CCRenderTexture.js            | not yet         |                                   |
|cocos2d/shaders/CCGLProgram.js                       | full            |                                   |
|cocos2d/shaders/CCGLStateCache.js                    | not yet         |                                   |
|cocos2d/shaders/CCShaderCache.js                     | not yet         |                                   |
|cocos2d/shaders/CCShaders.js                         | not yet         |                                   |
|cocos2d/shape-nodes/CCDrawNode.js                    | not yet         |                                   |
|cocos2d/text-input/CCIMEDispatcher.js                | not yet         |                                   |
|cocos2d/text-input/CCTextFieldTTF.js                 | not yet         |                                   |
|cocos2d/tilemap/CCTGAlib.js                          | not yet         |                                   |
|cocos2d/tilemap/CCTMXLayer.js                        | not yet         |                                   |
|cocos2d/tilemap/CCTMXObjectGroup.js                  | not yet         |                                   |
|cocos2d/tilemap/CCTMXTiledMap.js                     | not yet         |                                   |
|cocos2d/tilemap/CCTMXXMLParser.js                    | not yet         |                                   |
|cocos2d/transitions/CCTransition.js                  | not yet         |                                   |
|cocos2d/transitions/CCTransitionPageTurn.js          | not yet         |                                   |
|cocos2d/transitions/CCTransitionProgress.js          | not yet         |                                   |
